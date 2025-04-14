<template>
  <div class="grid grid-cols-3 gap-2 p-3 bg-base-300">
    <div class="">
      <h1 class="text-4xl">Meal Finder</h1>
    </div>
    <div>
      <h1><input class="input" type="text" placeholder="Search..." /></h1>
    </div>
    <div class="flex justify-end">
      <Icon class="mx-2" name="solar:book-2-outline" style="width: 2.5rem; height: 2.5rem;" />
      <Icon class="mx-2" name="solar:heart-bold" style="width: 2.5rem; height: 2.5rem;" />
      <Icon class="mx-2" name="solar:user-bold" style="width: 2.5rem; height: 2.5rem;" />
    </div>
  </div>

  <div class="flex justify-around py-2">
    <NuxtLink to="/" class="text-xl text-primary font-semibold">Europa</NuxtLink>
    <NuxtLink to="/login" class="text-xl text-secondary font-bold">Azja</NuxtLink>
    <NuxtLink to="/theme-test" class="text-xl text-accent font-extrabold">Afryka</NuxtLink>
    <NuxtLink to="/registration" class="text-xl">Ameryka Południowa </NuxtLink>
    <NuxtLink to="/registration" class="text-xl">Ameryka Północna </NuxtLink>
    <NuxtLink to="/registration" class="text-xl">Australia i Oceania</NuxtLink>
  </div>
</template>

<script lang="ts" setup>

</script>

<style></style>