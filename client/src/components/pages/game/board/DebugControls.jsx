const monopolyGroups = {
  brown: [1, 3],
  lightBlue: [6, 8, 9],
  pink: [11, 13, 14],
  orange: [16, 18, 19],
  red: [21, 23, 24],
  yellow: [26, 27, 29],
  green: [31, 32, 34],
  darkBlue: [37, 39], // Boardwalk & Park Place
};

import { useState } from "react";
import { useWebSocket } from "../../../../contexts/WebSocketProvider";
import { Button } from "../../../ui/button";

const DebugControls = () => {
  const { sendMessage } = useWebSocket();
  const [selectedGroup, setSelectedGroup] = useState("brown");

  return (
    <div className="mt-4 p-2 border border-red-500 rounded">
      <p className="text-sm text-red-600 mb-2">Debug Tools</p>
      
      <select
        value={selectedGroup}
        onChange={(e) => setSelectedGroup(e.target.value)}
        className="border p-1 mr-2"
      >
        {Object.keys(monopolyGroups).map((group) => (
          <option key={group} value={group}>
            {group}
          </option>
        ))}
      </select>

      <Button
        onClick={() => {
          monopolyGroups[selectedGroup].forEach((index) => {
            sendMessage("debug_give_property", index);
          });
        }}
      >
        Give Monopoly
      </Button>
    </div>
  );
};

export default DebugControls;