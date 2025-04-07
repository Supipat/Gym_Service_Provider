import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import "./requestapi.css";

export default function RequestAPI() {
  const navigate = useNavigate();

  const [apiKey, setApiKey] = useState("");
  const [gymId, setGymId] = useState("");
  const [search, setSearch] = useState("");
  const [sort, setSort] = useState("");
  const [responseData, setResponseData] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await axios.get(`http://localhost:8088/gymclass/${gymId}`, {
        headers: {
          Authorization: apiKey,
        },
        params: {
          search: search,
          sort: sort,
          order: "asc",
        },
      });
      console.log(res.data);
      setResponseData(res.data);
      navigate("/Affiliator", { state: { classes: res.data } });
    } catch (err) {
      console.log(err.responseM.dataM.error || err.message);
      alert("Error: " + (err.response?.data?.error || err.message));
    }
  };

  const goBack = () => {
    navigate("/Affiliator");
  };

  return (
    <div className="container-api">
      <div className="form-wrapper">
        <div className="header-row">
          <button className="back-button" onClick={goBack}>
            &#x2039;
          </button>
          <h1 className="heading">Request API</h1>
        </div>

        <form onSubmit={handleSubmit}>
          <div>
            <label htmlFor="apiKey">
              API Key<span className="required">*</span>
            </label>
            <input
              id="apiKey"
              type="text"
              placeholder="Enter your API key"
              required
              className="input"
              value={apiKey}
              onChange={(e) => setApiKey(e.target.value)}
            />
          </div>

          <div>
            <label htmlFor="gymId">
              Gym ID<span className="required">*</span>
            </label>
            <input
              id="gymId"
              type="number"
              placeholder="ใส่ Gym ID เช่น 1"
              required
              className="input"
              value={gymId}
              onChange={(e) => setGymId(e.target.value)}
            />
          </div>

          <div>
            <label htmlFor="search">Search</label>
            <input
              id="search"
              type="text"
              placeholder="คำค้นหา"
              className="input"
              value={search}
              onChange={(e) => setSearch(e.target.value)}
            />
          </div>

          <div>
            <label htmlFor="sort">Sort</label>
            <select
              id="sort"
              className="input"
              value={sort}
              onChange={(e) => setSort(e.target.value)}
            >
              <option value="">เลือกตัวเลือก</option>
              <option value="name">ชื่อคลาส</option>
              <option value="price">ราคา</option>
              <option value="level">ระดับ</option>
            </select>
          </div>
          <div className="button-container">
            <button type="submit" className="button">
              ค้นหา
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
