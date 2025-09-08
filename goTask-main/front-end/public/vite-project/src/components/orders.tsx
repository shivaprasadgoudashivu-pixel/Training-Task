import React, { useState } from "react";
import type { Order } from "../type";

export default function OrderBook() {
  const [orders, setOrders] = useState<Order[]>([]);
  const [loading, setLoading] = useState(false);
  const [userId, setUserId] = useState(""); // input field

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
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">Order Book</h1>

      {/* Search */}
      <div className="flex mb-4 gap-2">
        <input
          type="number"
          placeholder="Enter User ID"
          value={userId}
          onChange={(e) => setUserId(e.target.value)}
          className="border px-3 py-2 rounded-md w-40"
        />
        <button
          onClick={fetchOrders}
          className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600"
        >
          Search
        </button>
      </div>

      {loading ? (
        <div className="text-center mt-10">Loading...</div>
      ) : orders.length === 0 ? (
        <div>No orders found.</div>
      ) : (
        <div className="overflow-x-auto">
          <table className="min-w-full border border-gray-300">
            <thead className="bg-gray-100">
              <tr>
                <th className="border px-4 py-2">ID</th>
                <th className="border px-4 py-2">Scheme</th>
                <th className="border px-4 py-2">Amount</th>
                <th className="border px-4 py-2">Units</th>
                <th className="border px-4 py-2">Status</th>
                <th className="border px-4 py-2">Placed At</th>
                <th className="border px-4 py-2">Confirmed At</th>
                <th className="border px-4 py-2">Report</th>
              </tr>
            </thead>
            <tbody>
              {orders.map((order) => (
                <tr key={order.id} className="text-center">
                  <td className="border px-4 py-2">{order.id}</td>
                  <td className="border px-4 py-2">{order.Scheme}</td>
                  <td className="border px-4 py-2">{order.Amount}</td>
                  <td className="border px-4 py-2">{order.Units}</td>
                  <td className="border px-4 py-2">{order.Status}</td>
                  <td className="border px-4 py-2">
                    {new Date(order.Placed_at * 1000).toLocaleString()}
                  </td>
                  <td className="border px-4 py-2">
                    {new Date(order.confirm_at * 1000).toLocaleString()}
                  </td>
                  <td className="border px-4 py-2">
                    <a
                      href={order.Contact_Url}
                      target="_blank"
                      className="text-blue-500 underline"
                    >
                      Download
                    </a>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
