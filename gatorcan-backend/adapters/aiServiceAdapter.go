package adapters

import (
	"bytes"
	"encoding/json"
	"gatorcan-backend/DTOs"
	"gatorcan-backend/errors"
	"log"
	"net/http"
)

func GetRecommendedCourses(enrolled []int, interests []string, logger *log.Logger) ([]dtos.CourseRecommendationResponseDTO, error) {
	// Create request payload
	input := dtos.CourseRecommendationRequestDTO{EnrolledIDs: enrolled, Keywords: interests}
	jsonData, err := json.Marshal(input)
	if err != nil {
		logger.Printf("failed to marshal request: %v", err)
		return nil, errors.ErrMicroserviceError
	}

	// Send request to recommendation service
	resp, err := http.Post("http://localhost:8000/recommend", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Printf("failed to send request to recommendation service: %v", err)
		return nil, errors.ErrMicroserviceError
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusNotFound:
			return nil, errors.ErrMicroserviceNotFound
		case http.StatusBadRequest:
			return nil, errors.ErrMicroserviceError
		case http.StatusInternalServerError:
			return nil, errors.ErrMicroserviceError
		case http.StatusGatewayTimeout:
			return nil, errors.ErrMicroserviceTimeout
		default:
			return nil, errors.ErrMicroserviceError
		}
	}

	// Define struct that matches the JSON response format
	var responseData struct {
		Recommendations []struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			Tags  string `json:"tags"`
		} `json:"recommendations"`
	}

	// Decode the response
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		logger.Printf("failed to decode response from recommendation service: %v", err)
		return nil, errors.ErrMicroserviceError
	}

	// Convert to DTOs
	var recommendations []dtos.CourseRecommendationResponseDTO
	for _, rec := range responseData.Recommendations {
		recommendations = append(recommendations, dtos.CourseRecommendationResponseDTO{
			Id:    rec.ID,
			Title: rec.Title,
			Tags:  rec.Tags,
		})
	}

	return recommendations, nil
}

func GetSummary(text string, logger *log.Logger) (dtos.TextSummaryResponseDTO, error) {
	var summary dtos.TextSummaryResponseDTO
	input := dtos.TextSummaryRequestDTO{Text: text}
	jsonData, err := json.Marshal(input)
	if err != nil {
		logger.Printf("failed to marshal request: %v", err)
		return summary, errors.ErrMicroserviceError
	}

	resp, err := http.Post("http://localhost:8000/summarize", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Printf("failed to send request to recommendation service: %v", err)
		return summary, errors.ErrMicroserviceError
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return summary, errors.ErrMicroserviceError
	}

	var responseData struct {
		Summary string `json:"summary"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		logger.Printf("failed to decode response from recommendation service: %v", err)
		return summary, errors.ErrMicroserviceError
	}

	summary.Summary = responseData.Summary

	return summary, nil
}
