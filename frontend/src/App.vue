<template>
  <v-app id="inspire">
    <v-navigation-drawer
        v-model="drawer"
        app
    >
      <v-list class="relative pr-5" dense>
        <template v-for="(item, index) in items">
          <v-list-item
              v-if="item.to"
              :key="index"
              link
              exact
              :to="{ name: item.to }"
          >
            <v-list-item-icon>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>{{ item.name }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>

        </template>
      </v-list>

    </v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>Application</v-toolbar-title>
    </v-app-bar>

    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
export default {
  data: () => ({
    drawer: null,
    items: [
      {
        name: "Главная",
        to: "home",
        active: false,
        icon: "mdi-account-outline",
      },
      {
        name: "Хосты",
        to: "hosts",
        icon: "mdi-tag-outline",
      },
      {
        name: "Сети",
        to: "nets",
        icon: "mdi-tag-outline",
      },
    ],

  }),
}
</script>