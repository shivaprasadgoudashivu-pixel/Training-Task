// App.tsx
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import StockTradingApp from "./components/placeorder";
import OrderBook from "./components/orders";

function App() {
  return (
    <Router>
      <div className="min-h-screen grid place-items-center bg-gray-50">
        <Routes>
          <Route path="/" element={<StockTradingApp />} />
          <Route path="/orders/:id" element={<OrderBook />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
