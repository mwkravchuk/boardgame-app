import { Route, Routes } from "react-router-dom";
import routes from "./routes.jsx";

const RenderRoutes = () => {
  return (
    <div className="flex-1 content-center bg-blue-200">
      <Routes>
        {routes.map((route, i) => {
          return <Route key={i} path={route.path} element={route.element}/>
        })}
      </Routes>
    </div>
  );
};

export default RenderRoutes;