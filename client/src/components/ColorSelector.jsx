const availableColors = ['red', 'orange', 'yellow', 'green', 'blue', 'purple'];

const ColorSelector = ({ selectedColor, onSelect }) => {

  return (
    <div className="flex gap-1">
      {availableColors.map((color) => (
        <button
          key={color}
          style={{
            backgroundColor: color,
            width: "30px",
            height: "30px",
            border: selectedColor === color ? "solid 2px black" : "solid 1px gray",
            cursor: "pointer",
          }}
          onClick={() => onSelect(color)}
        ></button>
      ))}
    </div>
  );
};

export default ColorSelector;