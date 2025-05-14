# Project Name: GatorCan
  GatorCan : E-Learning Platform

## Visual Demo Links
- [Demo](https://drive.google.com/file/d/1LywboJhJ_LxYwHm6a4FTGKUGW-ZBsTMz/view?usp=sharing)
  
# Canvas Group Name : Code Mavericks


# Description:
    GatorCan is a full-stack educational platform designed to enhance the academic experience for students and instructors. 
    GatorCan allows for efficient course management, assignment submissions as well as intuitive collaboration.

# Core Features
    1. Enroll in or drop courses with ease.
    2. Access comprehensive course details, schedules, and materials.
    3. View a personalized weekly schedule for all enrolled courses.
    4. Submit assignments by uploading solutions directly.
    5. Instructors can review and manage submitted files efficiently.
    6. Engage in real-time discussions with instructors and peers.
    7. Dedicated messaging channels for each course.
   
# Technology Stack
    - Frontend: React.js for a dynamic and responsive user interface.
    - Backend: Golang for robust and scalable server-side operations.
    - Database: SQLite for lightweight and efficient data storage.
    - Cloud Services AWS
  
# Future Scope:
    - AI/NLP Integration:
      - Personalized course recommendations.
      - Summarization of course materials.
      - AI-powered chat assistance for queries.

# Team Members:
    # Frontend
      1. Harsh Gupta (gupta.harsh@ufl.edu)
      2. Navnit Krishna Pasupuleti (pasupuleti.n@ufl.edu)
      
    # Backend
      1. Muthukumaran Ulaganathan (ulaganathan.m@ufl.edu)
      2. Mohammad Wael (m.mohammadwael@ufl.edu)

# Steps to Run:
    Clone the repository
    # Backend
      1. cd gatorcan-backend
      2. go mod download
      3. go mod tidy
      4. go build
      5. go run main.go
    # AIService
      1. cd AIservice
      2. uvicorn main:app --reload --port 8000
    # Frontend
      1. cd gatorcan-frontend
      2. npm install
      3. npm run dev
