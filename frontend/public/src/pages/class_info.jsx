import React from "react";
import { Link } from "react-router-dom";

const ClassInfo = () => {
    return (
        <section className="text-center text-lg-start">
            <style>
                {`
          .cascading-right {
            margin-right: -50px;
          }

          @media (max-width: 991.98px) {
            .cascading-right {
              margin-right: 0;
            }
          }
        `}
            </style>

            {/* Jumbotron */}
            <div className="container py-4">
                <div className="row g-0 align-items-center">

                    {/*  Image Section moved to the left */}
                    <div className="col-lg-6 mb-5 mb-lg-0">
                        <img
                            src="img/cover.jpg"
                            className="w-100 rounded-4 shadow-4"
                            alt="Signup Illustration"
                        />
                    </div>

                    {/*  Form Section moved to the right */}
                    <div className="col-lg-6 mb-5 mb-lg-0">
                        <div
                            className="card cascading-right bg-body-tertiary"
                            style={{ backdropFilter: "blur(30px)" }}
                        >
                            <div className="card-body p-5 shadow-5 text-center">
                                <h2 className="fw-bold mb-5">Gym name</h2>
                                <form>
                                    {/* First & Last Name */}
                                    <div className="row mb-4">
                                        <div className="col-12 text-start">
                                            <p><strong>ชื่อคลาส:</strong> class name </p>
                                        </div>
                                        <div className="col-12 text-start">
                                            <p><strong>คำอธิบาย:</strong> description </p>
                                        </div>
                                        <div className="col-12 text-start">
                                            <p><strong>ประเภทคลาส:</strong> class type </p>
                                        </div>
                                        <div className="col-12 text-start">
                                            <p><strong>ราคา:</strong> price </p>
                                        </div>
                                        <div className="col-12 text-start">
                                            <p><strong>ระดับความยาก-ง่าย:</strong> easy medium  hard </p>
                                        </div>
                                        <div className="col-12 text-start">
                                            <p><strong>วันที่เริ่มเรียน:</strong> YY/MM/DD </p>
                                        </div>
                                        <div className="col-12 text-start">
                                            <p><strong>เวลาเรียน:</strong> 60 MIN/DAY </p>
                                        </div>
                                    </div>

                                    <Link to="/home"> /* รอแก้ */
                                        <button type="button" className="btn btn-primary btn-block mb-4">
                                            ย้อนกลับ
                                        </button>
                                    </Link>

                                </form>
                            </div>
                        </div>
                    </div>

                </div>
            </div>
            {/* Jumbotron */}
        </section>
    );
};

export default ClassInfo;
