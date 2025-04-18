from fastapi import FastAPI, Query
from pydantic import BaseModel
from typing import List
import pandas as pd
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity
import numpy as np
import re

app = FastAPI()

# Mock DB (replace with actual DB or file later)
courses = pd.DataFrame([
    {"id": 1, "title": "Intro to Machine Learning", "tags": "ML AI supervised unsupervised"},
    {"id": 2, "title": "Cloud Computing", "tags": "AWS GCP Azure DevOps"},
    {"id": 3, "title": "Natural Language Processing", "tags": "NLP text AI BERT"},
    {"id": 4, "title": "Database Systems", "tags": "SQL NoSQL DBMS indexing"},
    {"id": 5, "title": "Computer Vision", "tags": "AI images CNN OpenCV"}
])

# Precompute TF-IDF vectors
tfidf = TfidfVectorizer()
tfidf_matrix = tfidf.fit_transform(courses["tags"])

class Input(BaseModel):
    enrolled_ids: List[int]
    keywords: List[str]

class TextRequest(BaseModel):
    text: str
    sentences_count: int = 3

@app.post("/recommend")
def recommend(input: Input):
    enrolled = courses[courses["id"].isin(input.enrolled_ids)]
    if enrolled.empty:
        return {"recommendations": []}

    # Vector for enrolled courses
    enrolled_vec = tfidf.transform(enrolled["tags"])
    mean_vec = enrolled_vec.mean(axis=0)

    # Convert mean_vec from matrix to array
    mean_vec = np.asarray(mean_vec)

    # Filter out already enrolled
    remaining = courses[~courses["id"].isin(input.enrolled_ids)]
    if input.keywords:
        keyword_str = " ".join(input.keywords)
        remaining = remaining[remaining["tags"].str.contains("|".join(input.keywords), case=False)]

    if remaining.empty:
        return {"recommendations": []}

    remaining_vecs = tfidf.transform(remaining["tags"])
    print(cosine_similarity(mean_vec, remaining_vecs))
    similarities = cosine_similarity(mean_vec, remaining_vecs)[0]

    remaining["score"] = similarities
    top = remaining.sort_values("score", ascending=False).head(3)

    return {"recommendations": top[["id", "title", "tags"]].to_dict(orient="records")}


@app.post("/summarize")
async def smart_summarize(request: TextRequest):
    summary = smart_summarizer(request.text, request.sentences_count)
    return {"summary": summary}

def smart_summarizer(text, sentences_count=3):
    # Clean and split into sentences
    sentences = re.split(r'(?<=[.!?]) +', text)
    if len(sentences) <= sentences_count:
        return text

    # Vectorize the sentences
    vectorizer = TfidfVectorizer(stop_words='english')
    tfidf_matrix = vectorizer.fit_transform(sentences)

    # Calculate cosine similarity between sentences
    similarity_matrix = cosine_similarity(tfidf_matrix)

    # Sentence scores: sum of cosine similarities with other sentences
    sentence_scores = similarity_matrix.sum(axis=1)

    # Boost scores for sentences that have more unique important words
    important_words = vectorizer.get_feature_names_out()
    for idx, sentence in enumerate(sentences):
        word_count = sum(1 for word in important_words if word.lower() in sentence.lower())
        sentence_scores[idx] += 0.1 * word_count  # Small boost

    # Pick top sentences
    ranked_sentences_idx = np.argsort(sentence_scores)[-sentences_count:]
    ranked_sentences_idx = sorted(ranked_sentences_idx)

    # Build final summary
    summary = ' '.join([sentences[i] for i in ranked_sentences_idx])
    return summary.strip()