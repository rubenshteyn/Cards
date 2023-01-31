<template>
  <div v-if="isLoggedIn()" class="row wrapper d-flex flex-column overflow-hidden">
    <main-header/>
    <UserMenu/>
    <router-view />
  </div>
  <div v-else class="d-flex flex-column flex-column-fluid flex-lg-row auth-content">
    <router-view />
  </div>
</template>

<script>
import mainHeader from "@/components/mainHeader"
import UserMenu from "@/components/userMenu.vue";
import { mapGetters } from "vuex";

export default {
  name: 'app',
  components: {
    mainHeader,
    UserMenu
  },
  data() {
    return {
      url: process.env.VUE_APP_URL,
      title: process.env.VUE_APP_TITLE
    }
  },
  async beforeCreate() {
    await this.$store.dispatch("auth/loadUser")
  },
  methods: {
    isLoggedIn() {
      if (this.isLoaded) {
        let currentPath = this.$route.path
        let auth = false
        if (currentPath === "/login") {
          auth = true
        }
        if (this.loggedIn) {
          if (auth) {
            this.$router.push("/")
          }
          return true
        }
        if (!auth) {
          this.$router.push("/login")
        }
        return false
      }
    }
  },
  computed: {
    ...mapGetters({
        loggedIn: "auth/isLoggedIn",
        isLoaded: "auth/isLoaded"
    })
  }
}
</script>

<style>
.grecaptcha-badge {
  display: none !important;
}
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  /* display: none; <- Crashes Chrome on hover */
  -webkit-appearance: none;
  margin: 0; /* <-- Apparently some margin are still there even though it's hidden */
}
tbody tr:hover {
  background-color: rgb(230, 255, 241) !important;
}
a:hover, a:hover svg,
button:hover, button:hover svg {
  transition: all 0.2s ease-in-out;
  opacity: 0.8;
}
body {
  margin: 0;
  position: relative;
  overflow-x: hidden;
}

.auth-content {
    height: 100vh;
    background: url('./assets/media/background.jpg') 0 0 / cover no-repeat;
}

.vs--searchable .vs__dropdown-toggle {
  background-color: var(--kt-input-solid-bg);
  border-color: var(--kt-input-solid-bg);
  color: var(--kt-input-solid-color);
  min-height: calc(1.5em + 1.65rem + 2px);
  padding: 0.825rem 1.5rem;
  font-size: 1.15rem;
  border-radius: 0.625rem;
}

.vs--open .vs__dropdown-menu {
  border-color: var(--kt-input-solid-bg-focus) !important;
}

.vs__dropdown-toggle {

}

.vs__selected, .vs__actions, .vs__search, .vs__search:focus, .vs__selected-options, .vs__actions {
  border: unset !important;
  padding: 0 !important;
  margin: 0 !important;
}

.vs__search {
  color: var(--kt-text-gray-500) !important;
}

.vs__search, .vs__search:focus, .vs__selected {
  font-size: 1.15rem !important;
  font-weight: 500;
}
.vs__selected {
  color: var(--kt-input-solid-color) !important;
}
.vs__clear {
  display: flex;
}
.modal-height {
  height: 62vh;
}

label {
  margin: 0;
}
</style>
