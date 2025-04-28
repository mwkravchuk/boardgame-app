import { Route, Routes } from "react-router-dom";
import routes from "./routes.jsx";

import styles from "./RenderRoutes.module.css";

const RenderRoutes = () => {
  return (
    <div className={styles.routesPage}>
      <Routes>
        {routes.map((route, i) => {
          return <Route key={i} path={route.path} element={route.element}/>
        })}
      </Routes>
    </div>
  );
};

export default RenderRoutes;