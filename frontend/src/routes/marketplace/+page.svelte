<div class="flex space-x-4 ml-64 pl-4 pr-4">
  <!-- Left Column - Contains both Filter and Suggested Projects -->
  <div class="w-2/5 space-y-4 flex-shrink-0">
    <!-- Filter Panel -->
    <div class="card p-4">
      <h3 class="text-base font-semibold mb-3">Filters</h3>
      <form class="space-y-3">
        <div>
          <label class="block text-sm font-medium mb-1">Skills</label>
          <select
            class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
          >
            <option>All Skills</option>
            <option>React</option>
            <option>Node.js</option>
            <option>Python</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium mb-1">Difficulty</label>
          <input type="range" min="1" max="5" value="3" class="w-full" />
          <p class="text-xs text-gray-500 mt-1">Intermediate</p>
        </div>
        <div>
          <label class="block text-sm font-medium mb-1">Field</label>
          <div class="space-y-1">
            <label class="flex items-center text-sm">
              <input type="checkbox" class="mr-2" /> Web Development
            </label>
            <label class="flex items-center text-sm">
              <input type="checkbox" class="mr-2" /> Data Science
            </label>
          </div>
        </div>
        <button type="submit" class="btn w-full">Apply Filters</button>
      </form>
    </div>

    <!-- Suggested Projects -->
    <div class="card p-3">
      <h3 class="text-base font-semibold mb-2">Suggested Projects</h3>
      <div class="space-y-2">
        <div class="flex justify-between items-center p-2 bg-gray-100 rounded">
          <div>
            <p class="text-sm font-medium">Build a Chat App</p>
            <p class="text-xs text-gray-500">
              Skills: React, Node.js | 90% Match
            </p>
          </div>
          <button class="btn text-xs">Apply</button>
        </div>
        <div class="flex justify-between items-center p-2 bg-gray-100 rounded">
          <div>
            <p class="text-sm font-medium">E-commerce Backend</p>
            <p class="text-xs text-gray-500">
              Skills: Python, MongoDB | 85% Match
            </p>
          </div>
          <button class="btn text-xs">Apply</button>
        </div>
      </div>
    </div>
  </div>

  <!-- Project List -->
  <div class="w-3/5 space-y-3">
    <!-- Search Bar -->
    <div class="flex items-center space-x-2">
      <input
        type="text"
        class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
        placeholder="Search projects..."
      />
      <button class="btn">
        <svg
          class="w-6.5 h-6.5"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
          ></path>
        </svg>
      </button>
    </div>

    <!-- Project Cards (unchanged) -->
    <div class="card p-3">
      <a href="/marketplace/project-detail">
        <div class="flex justify-between items-center">
          <div>
            <h4 class="text-base font-medium">Build a Chat App</h4>
            <p class="text-sm text-gray-500">
              Skills: React, Node.js | 90% Match
            </p>
            <p class="text-xs text-gray-400 mt-1">
              Timeline: 2 weeks | Posted: 15/03/2025
            </p>
          </div>

          <div class="flex flex-row gap-2">
            <p class="text-sm text-gray-500">3/5</p>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="gray"
              height="20px"
              width="px"
              version="1.1"
              id="Layer_1"
              viewBox="0 0 315.539 315.539"
              xml:space="preserve"
            >
              <g>
                <g>
                  <g>
                    <path
                      d="M38.877,38.884C17.44,38.884,0,56.325,0,77.761c0,21.436,17.44,38.877,38.877,38.877s38.877-17.44,38.877-38.877     S60.314,38.884,38.877,38.884z M38.877,101.638C25.711,101.638,15,90.927,15,77.761s10.711-23.877,23.877-23.877     s23.877,10.711,23.877,23.877S52.043,101.638,38.877,101.638z"
                    />
                    <path
                      d="M38.877,131.65c-20.073,0-36.404,16.331-36.404,36.404v101.101c0,4.143,3.358,7.5,7.5,7.5h57.808     c4.142,0,7.5-3.357,7.5-7.5V168.054C75.281,147.981,58.95,131.65,38.877,131.65z M60.281,261.655H17.473v-93.601     c0-11.802,9.602-21.404,21.404-21.404c11.803,0,21.404,9.601,21.404,21.404V261.655z"
                    />
                    <path
                      d="M157.77,38.884c-21.437,0-38.877,17.44-38.877,38.877s17.44,38.877,38.877,38.877s38.877-17.44,38.877-38.877     S179.207,38.884,157.77,38.884z M157.77,101.638c-13.166,0-23.877-10.711-23.877-23.877s10.711-23.877,23.877-23.877     s23.877,10.711,23.877,23.877S170.936,101.638,157.77,101.638z"
                    />
                    <path
                      d="M157.769,131.65c-20.073,0-36.404,16.331-36.404,36.404v101.101c0,4.143,3.358,7.5,7.5,7.5h57.808     c4.142,0,7.5-3.357,7.5-7.5V168.054C194.173,147.981,177.842,131.65,157.769,131.65z M179.174,261.655h-42.808v-93.601     c0-11.802,9.602-21.404,21.404-21.404s21.404,9.601,21.404,21.404V261.655z"
                    />
                    <path
                      d="M276.662,116.638c21.437,0,38.877-17.44,38.877-38.877s-17.44-38.877-38.877-38.877s-38.877,17.44-38.877,38.877     S255.225,116.638,276.662,116.638z M276.662,53.884c13.166,0,23.877,10.711,23.877,23.877s-10.711,23.877-23.877,23.877     c-13.166,0-23.877-10.711-23.877-23.877S263.496,53.884,276.662,53.884z"
                    />
                    <path
                      d="M276.662,131.65c-20.073,0-36.404,16.331-36.404,36.404v101.101c0,4.143,3.358,7.5,7.5,7.5h57.808     c4.142,0,7.5-3.357,7.5-7.5V168.054C313.066,147.981,296.735,131.65,276.662,131.65z M298.066,261.655h-42.808v-93.601     c0-11.802,9.602-21.404,21.404-21.404s21.404,9.601,21.404,21.404V261.655z"
                    />
                  </g>
                </g>
              </g>
            </svg>
          </div>
        </div>
      </a>
    </div>

    <!-- Other project cards continue here -->
  </div>
</div>
