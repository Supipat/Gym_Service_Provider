import "../style.css";
import "../bootstrap.min.css";
import React, { useState } from "react";
import { Helmet } from "react-helmet";
import Navbar from "../component/navbar";
import { faPhone } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Footer from "../component/footer";

const Home = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  
  return (
    <div>
      <Helmet>
        <meta charset="utf-8" />
        <title>Gym Service Provider</title>
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <meta name="keywords" content="" />
        <meta name="description" content="" />

        {/* Google Fonts */}
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
        <link
          href="https://fonts.googleapis.com/css2?family=Open+Sans:wght@400;600&family=Raleway:wght@600;800&display=swap"
          rel="stylesheet"
        />

        {/* Font Awesome */}
        <link
          rel="stylesheet"
          href="https://use.fontawesome.com/releases/v5.15.4/css/all.css"
        />
        <link
          href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.4.1/font/bootstrap-icons.css"
          rel="stylesheet"
        />

        {/* Other Styles */}
        <link href="lib/lightbox/css/lightbox.min.css" rel="stylesheet" />
        <link
          href="lib/owlcarousel/assets/owl.carousel.min.css"
          rel="stylesheet"
        />
        <link href="css/bootstrap.min.css" rel="stylesheet" />
        <link href="css/style.css" rel="stylesheet" />
      </Helmet>

      <Navbar />
      <div class="container-fluid featurs py-5" >
        <div class="container py-5">
          <div class="row g-4 rounded bg-light p-4">
            <div class="featurs-item text-center rounded bg-light p-4">
              <div class="featurs-content text-center">
                <h5>เราคือ</h5>
                <p class="mb-0">
                  แพลตฟอร์ม Affiliate Service สำหรับการจองคลาสออกกำลังกาย
                  ที่ช่วยให้คุณแนะนำและสร้างรายได้จากการเชื่อมต่อลูกค้ากับยิมชั้นนำทั่วประเทศ!
                  จองง่าย ได้ค่าคอมทันที!
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="container-fluid service py-5" >
        <div className="container">
          <h1 class="mb-2">Gym ชั้นนำ</h1>

          <div className="row g-4 justify-content-center">
            <div className="col-md-6 col-lg-4">
              <a href={isLoggedIn ? "#" : "#login"} onClick={(e) => {
                if (!isLoggedIn) {
                  e.preventDefault();
                  alert("กรุณาล็อกอินหรือลงทะเบียนก่อนเข้าสู่บริการ");
                }
              }}>
                <div className="service-item bg-secondary rounded border border-secondary">
                <img src="img/fitness6.jpg" className="img-fluid rounded-top w-100" alt="" />
                  <div className="px-4 rounded-bottom">
                    <div className="service-content bg-primary text-center p-3 rounded">
                      <h5 className="text-white">FitZone Gym</h5>
                      <h6 className="mb-0">123 Sukhumvit Road, Bangkok</h6>
                      <div className="d-flex justify-content-center align-items-center mt-2">
                        <FontAwesomeIcon icon={faPhone} className="mb-0 text-white" style={{ marginRight: '8px' }} />
                        <h6 className="mb-0 text-white">098-123-4567</h6>
                      </div>
                    </div>
                  </div>
                </div>
              </a>
            </div>
            <div className="col-md-6 col-lg-4">
            <a href={isLoggedIn ? "#" : "#login"} onClick={(e) => {
                if (!isLoggedIn) {
                  e.preventDefault(); 
                  alert("กรุณาล็อกอินหรือลงทะเบียนก่อนเข้าสู่บริการ");
                }
              }}>
                <div className="service-item bg-dark rounded border border-dark">
                  <img src="img/boxing2.jpg" className="img-fluid rounded-top w-100" alt="" />
                  <div className="px-4 rounded-bottom">
                    <div className="service-content bg-light text-center p-4 rounded">
                      <h5 className="text-primary">Elite Boxing Club</h5>
                      <h6 className="mb-0">45 Silom Road, Bangkok</h6>
                      <div className="d-flex justify-content-center align-items-center mt-2">
                        <FontAwesomeIcon icon={faPhone} className="mb-0 text-primary" style={{ marginRight: '8px' }} />
                        <h6 className="mb-0 text-primary">089-987-6543</h6>
                      </div>
                    </div>
                  </div>
                </div>
              </a>
            </div>
            <div className="col-md-6 col-lg-4">
            <a href={isLoggedIn ? "#" : "#login"} onClick={(e) => {
                if (!isLoggedIn) {
                  e.preventDefault();
                  alert("กรุณาล็อกอินหรือลงทะเบียนก่อนเข้าสู่บริการ");
                }
              }}>
                <div className="service-item bg-primary rounded border border-primary">
                  <img src="img/yoga1.jpg" className="img-fluid rounded-top w-100" alt="" />
                  <div className="px-4 rounded-bottom">
                    <div className="service-content bg-secondary text-center p-3 rounded">
                      <h5 className="text-white">Flex Yoga Studio</h5>
                      <h6 className="mb-0">99 Rama IV Road, Bangkok</h6>
                      <div className="d-flex justify-content-center align-items-center mt-2">
                        <FontAwesomeIcon icon={faPhone} className="mb-0 text-white" style={{ marginRight: '8px' }} />
                        <h6 className="mb-0 text-white">092-345-6789</h6>
                      </div>
                    </div>
                  </div>
                </div>
              </a>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
};

export default Home;
