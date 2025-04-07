import axios from "axios";
import React from "react";
import { Link, useLocation } from "react-router-dom";

const Profile = () => {
  const location = useLocation();
  const classes = location.state?.classes;

  console.log(classes);

  const handleClick = async (classItem) => {
    try {
      const clickLog = {
        class_url: classItem.class_url,
        class_id: classItem.class_id,
        class_type: classItem.class_type,
        click_timestamp: new Date().toISOString(),
      };

      await axios.post("http://localhost:8088/api/clicklog", clickLog);

      window.open(classItem.class_url, "_blank");
      console.log(clickLog);
    } catch (error) {
      console.error("Error logging click:", error);
    }
  };

  return (
    <div>
      {/* Profile Header */}
      <div
        className="container py-3 px-4 rounded"
        style={{ marginTop: "50px", backgroundColor: "#e9ecef" }}
      >
        <div className="d-flex justify-content-between align-items-center">
          <div className="d-flex align-items-center">
            <img
              src="img/website.jpg"
              className="img-fluid rounded-circle me-5"
              alt=""
              style={{
                width: "100px",
                height: "100px",
                objectFit: "cover",
                marginLeft: "15px",
              }}
            />
            <h1 className="text-black">My Website</h1>
          </div>
          <div>
            <Link to="/request" className="btn btn-primary text-white">
              Add affiliate products
            </Link>
          </div>
        </div>
      </div>

      {/* Classes Section */}
      <div className="container contact py-5 px-4">
        <div className="p-5 rounded bg-light">
          <div className="row g-4">
            {classes?.length > 0 ? (
              classes.map((classItem) => (
                <div
                  key={classItem.class_id} // ใช้ class_id เป็น key
                  className="card mb-4 p-4 shadow-sm"
                  style={{ maxWidth: "1000px", width: "100%", padding: "10px" }}
                >
                  <div className="row g-0 align-items-center">
                    <div className="col-md-4 text-center">
                      <img
                        src="img/class.jpg"
                        className="img-fluid"
                        alt="Class"
                        style={{ width: "150px", height: "150px" }}
                      />
                    </div>
                    <div className="col-md-8">
                      <div className="d-flex flex-column justify-content-between h-100">
                        <div className="card-body p-0">
                          <h5 className="card-title">{classItem.class_name}</h5>
                          <p className="card-text mb-1">
                            เกี่ยวกับคลาส: {classItem.description}
                          </p>
                          <p className="card-text mb-1">
                            ประเภทคลาส: {classItem.class_type}
                          </p>
                          <p className="card-text mb-1">
                            เวลา: {classItem.class_duration} นาที
                          </p>
                          <p className="card-text mb-1">
                            วันที่: {classItem.class_schedule}
                          </p>
                          <p className="card-text mb-1">
                            ราคา: {classItem.class_price} บาท
                          </p>
                        </div>
                        <div className="d-flex justify-content-end mt-3">
                          <a
                            href={classItem.class_url}
                            target="_blank"
                            rel="noopener noreferrer"
                            className="btn btn-outline-primary"
                            onClick={() => handleClick(classItem)} // เมื่อคลิกให้เรียก function
                          >
                            More detail
                          </a>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              ))
            ) : (
              <p className="text-center">ยังไม่มีข้อมูลคลาส</p>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Profile;
