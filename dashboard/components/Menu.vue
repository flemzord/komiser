<template>
  <nav class="mt-10">
    <a
      class="
        flex
        items-center
        mt-4
        py-2
        px-6
        bg-gray-700 bg-opacity-25
        text-gray-100
      "
      href="/"
    >
      <font-awesome-icon :icon="['fas', 'house-damage']" />
      <span class="mx-3">Dashboard</span>
    </a>

    <a
      class="
        flex
        items-center
        mt-4
        py-2
        px-6
        text-gray-500
        hover:bg-gray-700 hover:bg-opacity-25 hover:text-gray-100
      "
      href="/"
    >
      <font-awesome-icon :icon="['fas', 'lock']" />
      <span class="mx-3">Compute</span>
    </a>

    <a
      class="
        flex
        items-center
        mt-4
        py-2
        px-6
        text-gray-500
        hover:bg-gray-700 hover:bg-opacity-25 hover:text-gray-100
      "
      ref="/"
    >
      <font-awesome-icon :icon="['fas', 'warehouse']" />
      <span class="mx-3">Storage</span>
    </a>

    <a
      class="
        flex
        items-center
        mt-4
        py-2
        px-6
        text-gray-500
        hover:bg-gray-700 hover:bg-opacity-25 hover:text-gray-100
      "
      ref="/"
    >
      <font-awesome-icon :icon="['fas', 'network-wired']" />
      <span class="mx-3">Network</span>
    </a>

    <a
      class="
        flex
        items-center
        mt-4
        py-2
        px-6
        text-gray-500
        hover:bg-gray-700 hover:bg-opacity-25 hover:text-gray-100
      "
      ref="/"
    >
      <font-awesome-icon :icon="['fas', 'shield-alt']" />
      <span class="mx-3">Security</span>
    </a>

    <a
      class="
        flex
        items-center
        mt-4
        py-2
        px-6
        text-gray-500
        hover:bg-gray-700 hover:bg-opacity-25 hover:text-gray-100
      "
      ref="/"
    >
      <font-awesome-icon :icon="['fas', 'database']" />
      <span class="mx-3">Data & AI</span>
    </a>
  </nav>
</template>
