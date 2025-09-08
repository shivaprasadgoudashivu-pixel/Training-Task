import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const StockTradingApp = () => {
  const [userId, setUserId] = useState<number>(0);
  const [schemeCode, setSchemeCode] = useState<string>("MF1");
  const [units, setUnits] = useState<number>(1);
  const [message, setMessage] = useState<string>("");
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const navigate = useNavigate()

  const handlePlaceOrder = async () => {
    if (userId <= 0) {
      setMessage("Please enter valid User ID");
      return;
    }

    if (units <= 0) {
      setMessage("Units must be greater than 0");
      return;
    }

    setIsLoading(true);
    setMessage("");

    try {
      const response = await fetch("http://localhost:8081/api/v1/orders", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          userId,
          units,
          scheme: "one time",
          schemeCode,
        }),
      });

      const responseData = await response.json();

      if (response.ok) {
        setMessage(`SuccessFully Order placed for ${schemeCode} (User ID: ${userId})`);
        // setUserId(0);
        // setSchemeCode("MF1");
        // setUnits(1);
      } else {
        let errorMsg =
          responseData.error ||
          responseData.message ||
          responseData.details ||
          "Unknown error occurred";
        setMessage(` ${errorMsg}`);
      }
    } catch (error) {
      console.error("Failed to place order:", error);
      setMessage(" Failed to connect to server. Please check your connection.");
    } finally {
      setIsLoading(false);
    }
  };

  // const orderBook = ()=>{
  //     // navigate(`/orders/${userId}`)
  // }

  return (
      <div className="w-full max-w-md bg-white rounded-2xl shadow-xl p-8 text-center">
        <h1 className="text-2xl font-extrabold text-gray-800 mb-8">
          Mutual Fund Order
        </h1>

        <div className="grid gap-6 text-left">
          {/* User ID */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              User ID
            </label>
            <input
              type="number"
              value={userId}
              onChange={(e) => setUserId(parseInt(e.target.value) || 0)}
              placeholder="Enter user ID"
              className="w-full border rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          {/* Scheme Code */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Scheme Code
            </label>
            <select
              value={schemeCode}
              onChange={(e) => setSchemeCode(e.target.value)}
              className="w-full border rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="MF1">MF1</option>
              <option value="MF2">MF2</option>
              <option value="MF3">MF3</option>
            </select>
          </div>

          {/* Units */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Units
            </label>
            <input
              type="number"
              onChange={(e) => setUnits(parseInt(e.target.value) || 1)}
              className="w-full border rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          {/* Place Order Button */}
          <button
            onClick={handlePlaceOrder}
            disabled={isLoading}
            className="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-3 rounded-xl shadow-md transition disabled:bg-blue-300 disabled:cursor-not-allowed"
          >
            {isLoading ? "Processing..." : "Place Order"}
          </button>
        </div>
        <br></br>
          <button
            onClick={()=>{
               navigate(`/orders/${userId}`)
            }}

            className="w-full bg-green-600 hover:bg-green-700 text-white font-semibold py-3 rounded-xl shadow-md transition"
          >
            View Orders
          </button>
        {/* Message */}
        {message && (
          <div
            className={`mt-6 p-4 rounded-xl font-medium text-center shadow-sm ${
              message.includes("SuccessFully")
                ? "bg-green-100 text-green-700"
                : "bg-red-100 text-red-700"
            }`}
          >
            {message}
          </div>
        )}
      </div>
  );
};

export default StockTradingApp;
