import { Repeat } from "lucide-react";

function Route({ start, end, onSelect }) {
  return (
    <button
      type="button"
      onClick={() => onSelect?.(start, end)}
      className="cursor-pointer flex flex-nowrap items-center justify-between gap-2 rounded-lg p-2 transition-colors hover:bg-gray-700/50"
    >
      <span className="min-w-0 flex-1 text-right text-sm font-semibold capitalize text-gray-200 whitespace-nowrap truncate">
        {start}
      </span>

      <Repeat className="text-gray-300" />

      <span className="min-w-0 flex-1 text-left text-sm font-semibold capitalize text-gray-200 whitespace-nowrap truncate">
        {end}
      </span>
    </button>
  );
}

export default Route;
