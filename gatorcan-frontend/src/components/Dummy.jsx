import React, { useState } from "react";
import { PutObjectCommand, ListObjectsV2Command } from "@aws-sdk/client-s3";
import s3Client from "../awsConfig";

const Dummy = () => {
  const [file, setFile] = useState(null);
  const [uploading, setUploading] = useState(false);
  const [message, setMessage] = useState("");

  const [objects, setObjects] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleFileChange = (event) => {
    setFile(event.target.files[0]);
  };

  const uploadFile = async () => {
    if (!file) {
      alert("Please select a file first!");
      return;
    }

    setUploading(true);
    setMessage("");

    // Convert file to an ArrayBuffer
    const arrayBuffer = await file.arrayBuffer();
    const fileBuffer = new Uint8Array(arrayBuffer); // Convert to Uint8Array

    const params = {
      Bucket: import.meta.env.VITE_S3_BUCKET_NAME,
      Key: file.name,
      Body: fileBuffer,
      ContentType: file.type,
    };

    try {
      const command = new PutObjectCommand(params);
      await s3Client.send(command);
      setMessage("Upload successful!");
    } catch (error) {
      setMessage(`Upload failed: ${error.message}`);
    } finally {
      setUploading(false);
    }
  };

  const fetchObjects = async () => {
    setLoading(true);
    setError("");
    try {
      const command = new ListObjectsV2Command({
        Bucket: import.meta.env.VITE_S3_BUCKET_NAME,
      });
      const response = await s3Client.send(command);
      setObjects(response.Contents || []);
    } catch (err) {
      setError(`Error: ${err.message}`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <div className="p-4 max-w-md mx-auto">
        <input type="file" onChange={handleFileChange} className="mb-2" />
        <button
          onClick={uploadFile}
          disabled={uploading}
          className="bg-blue-500 text-white p-2 rounded"
        >
          {uploading ? "Uploading..." : "Upload to S3"}
        </button>
        {message && <p className="mt-2">{message}</p>}
      </div>
      <div className="p-4 max-w-md mx-auto">
        <button
          onClick={fetchObjects}
          className="bg-blue-500 text-white p-2 rounded"
        >
          List S3 Objects
        </button>
        {loading && <p>Loading...</p>}
        {error && <p className="text-red-500">{error}</p>}
        <ul>
          {objects.map((obj, index) => (
            <li key={index}>
              {obj.Key} - {obj.Size} bytes
            </li>
          ))}
        </ul>
      </div>
    </>
  );
};

export default Dummy;
