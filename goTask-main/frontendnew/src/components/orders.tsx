import React, { useEffect, useState } from "react";
import type { Order } from "../type.ts";
import { useParams } from "react-router-dom";

export default function OrderBook() {
  const [orders, setOrders] = useState<Order[]>([]);
  const [loading, setLoading] = useState(false);




  const { id } = useParams();
const [userId, setUserId] = useState(id ?? "");
  React.useEffect(()=>{
  
    fetchOrders()
  },[userId]);


  const fetchOrders = async () => {
    if (!userId) return;
    setLoading(true);

    try {
      const res = await fetch(`http://localhost:8081/api/v1/orders/${userId}`);
      if (!res.ok) throw new Error("Failed to fetch orders");
      const data: Order[] = await res.json();
      setOrders(data);
    } catch (err) {
      console.error("Error fetching orders:", err);
      setOrders([]);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 p-6">
      <div className="max-w-7xl mx-auto">
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-800 mb-2">Order Book</h1>
        </div>

        {/* Search Section */}
        <div className="bg-white rounded-xl shadow-sm p-6 mb-8">
          <div className="flex flex-col sm:flex-row items-start sm:items-center gap-4">
            <div className="w-full sm:w-auto">
              <label htmlFor="userId" className="block text-sm font-medium text-gray-700 mb-1">
                User ID
              </label>
              <input
                id="userId"
                type="number"
                placeholder="Enter User ID"
                value={userId}
                onChange={(e) => setUserId(e.target.value)}
                className="border border-gray-300 rounded-lg px-4 py-2.5 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 w-full sm:w-64"
              />
            </div>
            {/* <button
              // onClick={fetchOrders}
              disabled={!userId || loading}
              className="bg-blue-600 text-white px-6 py-2.5 rounded-lg hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:bg-blue-400 disabled:cursor-not-allowed mt-6 sm:mt-6 w-full sm:w-auto"
            >
              {loading ? "Searching..." : "Search Orders"}
            </button> */}
          </div>
        </div>

        {/* Results Section */}
        {loading ? (
          <div className="flex justify-center items-center h-64">
            <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
          </div>
        ) : orders.length === 0 ? (
          <div className="bg-white rounded-xl shadow-sm p-8 text-center">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-16 w-16 mx-auto text-gray-400 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
            <h3 className="text-lg font-medium text-gray-900 mb-1">No orders found</h3>
            <p className="text-gray-500">Enter a User ID </p>
          </div>
        ) : (
          <div className="bg-white rounded-xl shadow-sm overflow-hidden">
            {/* Table Header */}
            <div className="hidden md:grid grid-cols-10 bg-gray-100 px-6 py-3 text-sm font-medium text-gray-700 uppercase tracking-wider">
              <div className="col-span-1">ID</div>
              <div className="col-span-1">Scheme</div>
              <div className="col-span-1">Amount</div>
              <div className="col-span-1">Units</div>
              <div className="col-span-1">Status</div>
              <div className="col-span-2">Placed At</div>
              <div className="col-span-2">Confirmed At</div>
              <div className="col-span-1">Report</div>
            </div>

            {/* Orders List */}
            <div className="divide-y divide-gray-200">
              {orders.map((order) => (
                <div key={order.id} className="grid grid-cols-1 md:grid-cols-10 gap-4 px-6 py-4 hover:bg-gray-50">
                  {/* ID */}
                  <div className="md:col-span-1">
                    <div className="text-xs text-gray-500 md:hidden">ID</div>
                    <div className="font-medium">{order.id}</div>
                  </div>
                  
                  {/* Scheme */}
                  <div className="md:col-span-1">
                    <div className="text-xs text-gray-500 md:hidden">Scheme</div>
                    <div>{order.Scheme}</div>
                  </div>
                  
                  {/* Amount */}
                  <div className="md:col-span-1">
                    <div className="text-xs text-gray-500 md:hidden">Amount</div>
                    <div>{order.Amount}</div>
                  </div>
                  
                  {/* Units */}
                  <div className="md:col-span-1">
                    <div className="text-xs text-gray-500 md:hidden">Units</div>
                    <div>{order.Units}</div>
                  </div>
                  
                  {/* Status */}
                  <div className="md:col-span-1">
                    <div className="text-xs text-gray-500 md:hidden">Status</div>
                    <span className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                      order.Status === 'completed' ? 'bg-green-100 text-green-800' :
                      order.Status === 'pending' ? 'bg-yellow-100 text-yellow-800' :
                      'bg-gray-100 text-gray-800'
                    }`}>
                      {order.Status}
                    </span>
                  </div>
                  
                  {/* Placed At */}
                  <div className="md:col-span-2">
                    <div className="text-xs text-gray-500 md:hidden">Placed At</div>
                    <div>{new Date(order.Placed_at * 1000).toLocaleString()}</div>
                  </div>
                  
                  {/* Confirmed At */}
                  <div className="md:col-span-2">
                    <div className="text-xs text-gray-500 md:hidden">Confirmed At</div>
                    <div>{new Date(order.confirm_at * 1000).toLocaleString()}</div>
                  </div>
                  
                  {/* Report */}
                  <div className="md:col-span-1">
                    <div className="text-xs text-gray-500 md:hidden">Report</div>
                    <a
                      href={order.Contact_Url}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="inline-flex items-center text-blue-600 hover:text-blue-800"
                    >
                      {/* <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                      </svg> */}
                      Download
                    </a>
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  );
}