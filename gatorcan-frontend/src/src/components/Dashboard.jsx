import MediaCard from "./CourseCard";
import Navbar from "./Navbar";

function Dashboard() {

  const courses = [
    [
      "CAP5771 - Intro to Data Science",
      "CAP5771 - Intro to Data Science CAP5771 Spring 2025",
      "forestgreen",
    ],
    [
      "CEN5035 - Software Engineering",
      "CEN5035 - Software Engineering CEN5035 Spring 2025",
      "darkorchid",
    ],
    [
      "COP5556 - Program Language Principles",
      "COP5556 - Program Language Principles COP5556 Spring 2025",
      "MediumVioletRed",
    ],
    [
      "CAP5771 - Intro to Data Science",
      "CAP5771 - Intro to Data Science CAP5771 Spring 2025",
      "gold",
    ],
  ];

  return (
    <>
      <Navbar />
      <div style={{ marginLeft: "120px" }}>
        <h1>Dashboard</h1>
        <hr />
        <br></br>
        <div
          style={{ display: "flex", flexWrap: "wrap", justifyContent: "left", flexDirection: "row", width:"80%" }}
        >
          {courses.map((course, index) => {
            return (
              <MediaCard
                key={index}
                text1={course[0]}
                text2={course[1]}
                color={course[2]}
              ></MediaCard>
            );
          })}
        </div>
      </div>
    </>
  );
}

export default Dashboard;
