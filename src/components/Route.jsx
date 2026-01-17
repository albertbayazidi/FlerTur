function Route({ start, end }) {
  return (
    <div className="flex enter justify-between gap-2 rounded-lg p-2 transition-colors hover:bg-gray-700/50">
      <span className="flex-1 text-right text-sm font-semibold capitalize text-gray-200">
        {start}
      </span>

      <span className="flex-shrink-0 text-white">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={1.5}
          stroke="currentColor"
          className="fg-white h-5 w-5"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M7.5 21 3 16.5m0 0 4.5-4.5M3 16.5h13.5m0-13.5L21 7.5m0 0-4.5 4.5M21 7.5H7.5"
          />
        </svg>
      </span>

      <span className="flex-1 text-left text-sm font-semibold capitalize text-gray-200">
        {end}
      </span>
    </div>
  );
}

export default Route;
