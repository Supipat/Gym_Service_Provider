import React from "react";
import NavbarRegis from "../component/navbarRegis";
import "../style.css";
import "../bootstrap.min.css";
import "./gym.css";
import "./goback.css";
import { useNavigate } from "react-router-dom";
import Footer from "../component/footer";

const Gym3 = () => {
  const navigate = useNavigate();
  const goBack = () => {
    navigate(-1);
  };

  return (
    <div>
      <NavbarRegis />
      <div
        style={{
          marginTop: "150px",
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          flexDirection: "column",
        }}
      >
        <div className="col-lg-7" style={{ marginTop: "30px" }}>
      
        <button className="back-button" onClick={goBack} style={{ marginBottom: "10px " }}>&#x2039;</button>
     
          <h2 >GET คลาสทั้งหมดของ Flex Yoga Studio</h2>
          
          <h4 style={{ marginTop: "15px" }}>คำอธิบาย</h4>          <p>
            Flex Yoga Studio เป็นคลาสออกกำลังกายที่ผสมผสานหลายศาสตร์
            เพื่อเสริมความแข็งแรง ยืดหยุ่น สมดุลของร่างกาย
            พร้อมช่วยลดความเครียดและผ่อนคลายจิตใจ
            เหมาะสำหรับผู้ที่ต้องการดูแลทั้งร่างกายและจิตใจอย่างสมดุล
          </p>
          <h4>Request</h4>
          <div
            style={{
              display: "grid",
              placeItems: "center",
              backgroundColor: "#f5f5f5",
              border: "1px solid #ccc",
              padding: "5px",
              borderRadius: "8px",
              fontFamily: "monospace",
              whiteSpace: "pre-wrap",
              wordWrap: "break-word",
              minHeight: "50px",
              marginBottom: "10px",
              marginTop: "10px",
            }}
          >
            <p style={{ margin: 0 }}>
              <span style={{ fontWeight: "bold", color: "green" }}>GET</span>{" "}
              http://localhost:8088/gymclass/3
            </p>
          </div>

          <h5>Data Dictionary</h5>
          <table className="table">
            <thead>
              <tr>
                <th scope="col">ชื่อ Field</th>
                <th scope="col">รายละเอียด</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>gymclass/3</td>
                <td>ยิม Flex Yoga Studio</td>
              </tr>

              <tr>
                <td>class_id</td>
                <td>รหัสคลาส</td>
              </tr>
              <tr>
                <td>class_name</td>
                <td>ชื่อคลาส</td>
              </tr>
              <tr>
                <td>description</td>
                <td>คำอธิบายรายละเอียดของคลาส</td>
              </tr>
              <tr>
                <td>class_type</td>
                <td>ประเภทของคลาส</td>
              </tr>
              <tr>
                <td>class_duration</td>
                <td>ระยะเวลาที่เรียน</td>
              </tr>
              <tr>
                <td>class_schedule</td>
                <td>วันที่เริ่มเรียน</td>
              </tr>
              <tr>
                <td>class_price</td>
                <td>ราคาของคลาส</td>
              </tr>
              <tr>
                <td>class_level</td>
                <td>ระดับความยากง่ายของคลาส</td>
              </tr>
            </tbody>
          </table>

          <h5 style={{ marginTop: "20px" }}>Response: 200 OK</h5>
          <div
            style={{
              backgroundColor: "#f5f5f5",
              border: "1px solid #ccc",
              padding: "15px",
              borderRadius: "8px",
              fontFamily: "monospace",
              whiteSpace: "pre-wrap",
              wordWrap: "break-word",
              maxHeight: "400px",
              overflowY: "auto",
            }}
          >
            <pre>
              {JSON.stringify(
                {
                  class_id: 7,
                  class_name: "Body Pump",
                  description:
                    "ยกน้ำหนักที่ผสมผสานกับการเคลื่อนไหวตามจังหวะเพลง เสริมสร้างกล้ามเนื้อและความแข็งแรงทั่วร่างกาย",
                  class_type: "Strength and Conditioning",
                  class_duration: "60",
                  class_schedule: "2025-05-03",
                  class_price: 300,
                  class_level: "ปานกลาง",
                },
                null,
                2
              )}
            </pre>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
};

export default Gym3;