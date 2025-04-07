import { faGlobe } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import NavbarRegis from "../component/navbarRegis";


import {
  faEnvelope,
  faEye,
  faEyeSlash,
  faKey,
  faPhone,
  faUser,
} from "@fortawesome/free-solid-svg-icons";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import Footer from "../component/footer";
import "./goback.css";

const affUrl = "http://localhost:8088/api/affiliator";

const User = () => {
  const [isHidden, setIsHidden] = useState(true);
  const [clientData, setClientData] = useState(null);

  const navigate = useNavigate();
  const goBack = () => {
    navigate(-1);
  };

  useEffect(() => {
    const fetchClient = async () => {
      const email = localStorage.getItem("user_email"); // ดึง email จาก localStorage

      if (!email) {
        alert("กรุณาล็อกอินก่อน");
        return;
      }

      try {
        const res = await axios.get(`${affUrl}/${email}`);
        setClientData(res.data); // คาดว่า backend จะ return ข้อมูลเฉพาะคนนี้
      } catch (err) {
        console.error("ดึงข้อมูลล้มเหลว:", err);
      }
    };

    fetchClient();
  }, []);
  return (
    <div>
      <NavbarRegis />
      <div
        className="container-fluid contact py-5"
        style={{ marginTop: "80px" }}
      >
        <div className="container py-5">
          <div className="p-5 bg-light rounded">
            <button className="back-button" onClick={goBack}>&#x2039;</button>
            <div className="row g-4">
              <div className="col-12">
                <div className="text-center" style={{ maxWidth: "100%" }}>
                  <h1 className="text-primary">Profile</h1>
                </div>
              </div>

              {/* Contact Info */}
              <div className="col-lg-5">
                <div className="d-flex p-4 rounded mb-4 bg-white">
                  <FontAwesomeIcon
                    icon={faUser}
                    size="2x"
                    className="text-primary me-4"
                  />
                  <div>
                    <h4>Name</h4>
                    <p className="mb-2">
                      {clientData?.affiliator_fname}{" "}
                      {clientData?.affiliator_lname}
                    </p>
                  </div>
                </div>
                <div className="d-flex p-4 rounded mb-4 bg-white">
                  <FontAwesomeIcon
                    icon={faEnvelope}
                    size="2x"
                    className="text-primary me-4"
                  />
                  <div>
                    <h4>Email</h4>
                    <p className="mb-2">{clientData?.affiliator_email}</p>
                  </div>
                </div>
                <div className="d-flex p-4 rounded bg-white">
                  <FontAwesomeIcon
                    icon={faPhone}
                    size="2x "
                    className="text-primary me-4"
                  />
                  <div>
                    <h4>Telephone</h4>
                    <p className="mb-2">{clientData?.affiliator_phone}</p>
                  </div>
                </div>
                <div className="d-flex p-4 rounded bg-white" style={{ marginTop: "25px" }}>
                <FontAwesomeIcon icon={faGlobe} size="2x" className="text-primary me-4" />

                  <div>
                    <h4>Websites</h4>
                    {clientData?.affiliator_website?.map((url, i) => (
                      <p key={i}>{url}</p>
                    ))}
                    <div>
                    <Link
                        to="/Affiliator"
                        target="_blank"
                            rel="noopener noreferrer"
                        className="btn btn-secondary text-brown"
                      >
                        ไปสู่หน้า Affiliator
                      </Link>
                    </div>
                  </div>
                </div>
              </div>

              <div className="col-lg-7">
                <div
                  className="d-flex p-4 rounded mb-4 bg-white position-relative"
                  style={{ height: "100%" }}
                >
                  <div style={{ flex: 1 }}>
                    <h4>
                      Api Token Key{" "}
                      <FontAwesomeIcon
                        icon={faKey}
                        className="text-primary me-2"
                      />
                    </h4>
                    <p
                      className="mb-2"
                      style={{
                        wordBreak: "break-word",
                        overflowWrap: "break-word",
                        maxWidth: "100%",
                        overflow: "hidden",
                        textOverflow: "ellipsis",
                      }}
                    >
                      {isHidden
                        ? "*************************"
                        : clientData?.api_key}
                    </p>
                  </div>

                  <FontAwesomeIcon
                    icon={isHidden ? faEye : faEyeSlash}
                    className="text-secondary"
                    style={{
                      position: "absolute",
                      right: "15px",

                      transform: "translateY(-50%)",
                      cursor: "pointer",
                    }}
                    onClick={() => setIsHidden(!isHidden)}
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </div>

  );
};

export default User;