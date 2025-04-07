import React from "react";
import { Helmet } from "react-helmet";
import Navbar from "../component/navbar";
import "../style.css";
import "../bootstrap.min.css";

const ApiList = () => {
    return (
        <>
            <Helmet>
                <title>API List</title>
            </Helmet>

            {/* Navbar */}
            <Navbar />
            <div>

            <h1 className="mb-4">Api</h1>

            <section className="intro" style={{ marginTop: "200px" }}>
            <div className="container">
                            <div className="row justify-content-center">
                                <div className="col-12">
                                    <div className="card">
                                        <div className="card-body">
                                            <div className="table-responsive">
                                                <table className="table table-hover mb-0">
                                                    <thead>
                                                        <tr>
                                                            <th scope="col"></th>
                                                            <th scope="col">Product Detail Views</th>
                                                            <th scope="col">Unique Purchases</th>
                                                            <th scope="col">Quantity</th>
                                                            <th scope="col">Product Revenue</th>
                                                            <th scope="col">Avg. Price</th>
                                                        </tr>
                                                    </thead>
                                                    <tbody>
                                                        <tr>
                                                            <th scope="row">Value</th>
                                                            <td>18,492</td>
                                                            <td>228</td>
                                                            <td>350</td>
                                                            <td>$4,787.64</td>
                                                            <td>$13.68</td>
                                                        </tr>
                                                        <tr>
                                                            <th scope="row">Percentage change</th>
                                                            <td className="text-danger">
                                                                <i className="fas fa-caret-down me-1"></i> -48.8%
                                                            </td>
                                                            <td className="text-success">
                                                                <i className="fas fa-caret-up me-1"></i> 14.0%
                                                            </td>
                                                            <td className="text-success">
                                                                <i className="fas fa-caret-up me-1"></i> 46.4%
                                                            </td>
                                                            <td className="text-success">
                                                                <i className="fas fa-caret-up me-1"></i> 29.6%
                                                            </td>
                                                            <td className="text-danger">
                                                                <i className="fas fa-caret-down me-1"></i> -11.5%
                                                            </td>
                                                        </tr>
                                                        <tr>
                                                            <th scope="row">Average</th>
                                                            <td className="text-danger">
                                                                <i className="fas fa-caret-down me-1"></i> -17,654
                                                            </td>
                                                            <td className="text-success">
                                                                <i className="fas fa-caret-up me-1"></i> 28
                                                            </td>
                                                            <td className="text-success">
                                                                <i className="fas fa-caret-up me-1"></i> 111
                                                            </td>
                                                            <td className="text-success">
                                                                <i className="fas fa-caret-up me-1"></i> $1,092.72
                                                            </td>
                                                            <td className="text-danger">
                                                                <i className="fas fa-caret-down me-1"></i> $-1.78
                                                            </td>
                                                        </tr>
                                                    </tbody>
                                                </table>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
            </section>
            </div>
        </>
    );
};

export default ApiList;
