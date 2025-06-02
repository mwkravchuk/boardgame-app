import Room from "../pages/Room";
import Lobby from "../pages/Lobby";
import Game from "../pages/game/Game";

const routes = [
  { path: "/", element: <Room /> },
  { path: "/lobby", element: <Lobby />},
  { path: "/game", element: <Game />}
];

export default routes;