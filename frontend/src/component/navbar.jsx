import { Link } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faUser, faBars, faUserFriends, faEnvelope } from "@fortawesome/free-solid-svg-icons";
import { Helmet } from "react-helmet";
import "../style.css";
import "../bootstrap.min.css";
import React from "react";

const Navbar = () => {
    return (
        <div className="container-fluid fixed-top">
            <Helmet>
                <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css" />
            </Helmet>

            {/* Top Bar */}
            <div className="container topbar bg-primary d-none d-lg-block">
                <div className="d-flex justify-content-between">
                    <div className="top-info ps-2">
                        <small className="me-3">
                            <FontAwesomeIcon icon={faUserFriends} className="me-2 text-secondary" />
                            <Link to="/" className="text-white">Notepads</Link>
                        </small>

                    </div>
                </div>
            </div>

            {/* Navbar */}
            <div className="container px-0">
                <nav className="navbar navbar-light bg-white navbar-expand-xl ">
                    <Link to="/" className="navbar-brand">
                        <h1 className="text-primary display-6">GymService</h1>
                    </Link>
                    <button className="navbar-toggler py-2 px-3" type="button" data-bs-toggle="collapse" data-bs-target="#navbarCollapse">
                        <FontAwesomeIcon icon={faBars} className="text-primary" />
                    </button>
                    <div className="collapse navbar-collapse bg-white" id="navbarCollapse">
                        <div className="navbar-nav mx-auto">
                        
                        </div>
                        <div className="d-flex m-3 me-0 align-items-center gap-4">
                            <Link to="/Register" className="my-auto">
                                <FontAwesomeIcon icon={faUser} size="2x" />
                            </Link>
                   
                        </div>
                    </div>
                </nav>
            </div>
        </div>
    );
};

export default Navbar;
