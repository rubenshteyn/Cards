<template>
  <div class="logged-in-user position-fixed d-flex align-items-center" style="z-index: 9999">
    <div class="text-dark mr-3">You are logged in as a user {{ currentUser()[0] }}</div>
    <div>
      <button class="btn text-dark btn-outline-light btn-sm" @click="returnToAdmin">Return</button>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
export default {
  name: "loggedInUser",
  methods: {
    ...mapActions({
      logout: "auth/logout",
    }),
    currentUser() {
      const user = this.$store.getters["auth/currentUser"];
      if (user) {
        return [user.email, user.access];
      }
      return "";
    },
    returnToAdmin() {
      this.logout()
      this.$router.push("/admin")
    }
  },
  computed: {
    ...mapGetters({
      user: "auth/currentUser"
    })
  }
}
</script>

<style scoped>
  .logged-in-user {
    border-radius: 10px;
    width: fit-content;
    background: rgba(208, 206, 206, 0.6);
    padding: 10px 15px;
    left: 0;
    right: 0;
    margin: 0 auto;
    bottom: 3rem;
  }
  .btn-outline-light {
    padding: 5px 10px !important;
    border: 1px solid #181C32 !important;
  }

</style>