<template>
  <div class="flex h-screen bg-gray-200">
    <div
      class="
        fixed
        z-30
        inset-y-0
        left-0
        w-64
        transition
        duration-300
        transform
        bg-gray-900
        overflow-y-auto
        lg:translate-x-0 lg:static lg:inset-0
      "
    >
      <div class="flex items-center justify-center mt-8">
        <div class="flex items-center">
          <div class="h-12 w-12">
            <Logo />
          </div>
          <span class="text-white text-2xl mx-2 font-semibold">Komiser</span>
        </div>
      </div>

      <Menu />
    </div>
    <div class="flex-1 flex flex-col overflow-hidden">
      <header
        class="
          flex
          justify-between
          items-center
          py-4
          px-6
          bg-white
          border-b-4 border-indigo-600
        "
      >
        <div class="flex items-center">
          <button class="text-gray-500 focus:outline-none lg:hidden">
            <svg
              class="h-6 w-6"
              viewBox="0 0 24 24"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M4 6H20M4 12H20M4 18H11"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
              ></path>
            </svg>
          </button>
        </div>

        <div class="flex items-center">
          <div class="relative">
            <button class="flex mx-4 text-gray-600 focus:outline-none">
              <svg
                class="h-6 w-6"
                viewBox="0 0 24 24"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M15 17H20L18.5951 15.5951C18.2141 15.2141 18 14.6973 18 14.1585V11C18 8.38757 16.3304 6.16509 14 5.34142V5C14 3.89543 13.1046 3 12 3C10.8954 3 10 3.89543 10 5V5.34142C7.66962 6.16509 6 8.38757 6 11V14.1585C6 14.6973 5.78595 15.2141 5.40493 15.5951L4 17H9M15 17V18C15 19.6569 13.6569 21 12 21C10.3431 21 9 19.6569 9 18V17M15 17H9"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                ></path>
              </svg>
            </button>
          </div>
        </div>
      </header>
      <Nuxt />
    </div>
  </div>
</template>
