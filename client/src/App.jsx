import { BrowserRouter } from "react-router-dom";

import Header from "./components/structure/Header";
import RenderRoutes from "./components/structure/RenderRoutes";
import Footer from "./components/structure/Footer";

const App = () => {
  return (
    <>
      <BrowserRouter>
        <div className="flex flex-col">
          {/* Make header and whatever route take up whole page (100vh) */}
          <div className="flex flex-col min-h-screen bg-amber-200">
            <div className="flex flex-1 justify-center items-center">
              <RenderRoutes />
            </div>
          </div> 
          <Footer />
        </div>
      </BrowserRouter>
    </>
  );
};

export default App;
