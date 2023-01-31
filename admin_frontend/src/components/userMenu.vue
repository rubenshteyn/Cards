<template>
  <div class="col">
    <!-- из шаблона start -->
    <div id="kt_aside" class="aside card" data-kt-drawer="true" data-kt-drawer-name="aside"
         data-kt-drawer-activate="{default: true, lg: false}" data-kt-drawer-overlay="true"
         data-kt-drawer-width="{default:'200px', '300px': '250px'}" data-kt-drawer-direction="start"
         data-kt-drawer-toggle="#kt_aside_toggle">
      <div class="aside-menu flex-column-fluid px-4">
        <!--begin::Aside Menu-->
        <div class="hover-scroll-overlay-y my-5 pe-4 me-n4" id="kt_aside_menu_wrapper" data-kt-scroll="true"
             data-kt-scroll-activate="true" data-kt-scroll-height="auto"
             data-kt-scroll-dependencies="{default: '#kt_aside_footer', lg: '#kt_header, #kt_aside_footer'}"
             data-kt-scroll-wrappers="#kt_aside, #kt_aside_menu" data-kt-scroll-offset="{default: '5px', lg: '75px'}"
             style="height: 678px;">
          <!--begin::Menu-->
          <div class="menu menu-column menu-rounded menu-sub-indention fw-semibold fs-6" id="#kt_aside_menu"
               data-kt-menu="true">
            <!--begin:Menu item-->
            <div data-kt-menu-trigger="click" class="menu-item menu-accordion show">
              <!--begin:Menu sub-->
              <!--begin:Menu item-->
              <div class="menu-item">
                <router-link @click="isShow = false" to="/" class="menu-link">
												<span class="menu-icon">
													<!--begin::Svg Icon | path: icons/duotune/general/gen025.svg-->
													<span class="svg-icon svg-icon-2 text-gray-400">
														<svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                 xmlns="http://www.w3.org/2000/svg">
															<rect x="2" y="2" width="9" height="9" rx="2" fill="currentColor"></rect>
															<rect opacity="0.3" x="13" y="2" width="9" height="9" rx="2" fill="currentColor"></rect>
															<rect opacity="0.3" x="13" y="13" width="9" height="9" rx="2" fill="currentColor"></rect>
															<rect opacity="0.3" x="2" y="13" width="9" height="9" rx="2" fill="currentColor"></rect>
														</svg>
													</span>
                          <!--end::Svg Icon-->
												</span>
                  <span class="menu-title fw-semibold">Dashboard</span>
                </router-link>
              </div>
              <!--end:Menu item-->
              <!--begin:Menu item-->
              <div class="menu-item">
                <router-link @click="isShow = false" to="/cards" class="menu-link">
												<span class="menu-icon">
													<!--begin::Svg Icon | path: icons/duotune/general/gen025.svg-->
													<span class="svg-icon svg-icon-2 text-gray-400">
                            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                                 class="bi bi-credit-card" viewBox="0 0 16 16">
                              <path
                                  d="M0 4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V4zm2-1a1 1 0 0 0-1 1v1h14V4a1 1 0 0 0-1-1H2zm13 4H1v5a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V7z"/>
                              <path d="M2 10a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-1z"/>
                            </svg>
													</span>
                          <!--end::Svg Icon-->
												</span>
                  <span class="menu-title fw-semibold">Cards List</span>
                </router-link>
              </div>
              <!--end:Menu item-->
              <!--begin:Menu item-->
              <div class="menu-item">
                <router-link @click="isShow = false" to="/transactions" class="menu-link">
												<span class="menu-icon">
													<!--begin::Svg Icon | path: icons/duotune/general/gen025.svg-->
													<span class="svg-icon svg-icon-2 text-gray-400">
                            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                                 class="bi bi-receipt" viewBox="0 0 16 16">
                              <path
                                  d="M1.92.506a.5.5 0 0 1 .434.14L3 1.293l.646-.647a.5.5 0 0 1 .708 0L5 1.293l.646-.647a.5.5 0 0 1 .708 0L7 1.293l.646-.647a.5.5 0 0 1 .708 0L9 1.293l.646-.647a.5.5 0 0 1 .708 0l.646.647.646-.647a.5.5 0 0 1 .708 0l.646.647.646-.647a.5.5 0 0 1 .801.13l.5 1A.5.5 0 0 1 15 2v12a.5.5 0 0 1-.053.224l-.5 1a.5.5 0 0 1-.8.13L13 14.707l-.646.647a.5.5 0 0 1-.708 0L11 14.707l-.646.647a.5.5 0 0 1-.708 0L9 14.707l-.646.647a.5.5 0 0 1-.708 0L7 14.707l-.646.647a.5.5 0 0 1-.708 0L5 14.707l-.646.647a.5.5 0 0 1-.708 0L3 14.707l-.646.647a.5.5 0 0 1-.801-.13l-.5-1A.5.5 0 0 1 1 14V2a.5.5 0 0 1 .053-.224l.5-1a.5.5 0 0 1 .367-.27zm.217 1.338L2 2.118v11.764l.137.274.51-.51a.5.5 0 0 1 .707 0l.646.647.646-.646a.5.5 0 0 1 .708 0l.646.646.646-.646a.5.5 0 0 1 .708 0l.646.646.646-.646a.5.5 0 0 1 .708 0l.646.646.646-.646a.5.5 0 0 1 .708 0l.646.646.646-.646a.5.5 0 0 1 .708 0l.509.509.137-.274V2.118l-.137-.274-.51.51a.5.5 0 0 1-.707 0L12 1.707l-.646.647a.5.5 0 0 1-.708 0L10 1.707l-.646.647a.5.5 0 0 1-.708 0L8 1.707l-.646.647a.5.5 0 0 1-.708 0L6 1.707l-.646.647a.5.5 0 0 1-.708 0L4 1.707l-.646.647a.5.5 0 0 1-.708 0l-.509-.51z"/>
                              <path
                                  d="M3 4.5a.5.5 0 0 1 .5-.5h6a.5.5 0 1 1 0 1h-6a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h6a.5.5 0 1 1 0 1h-6a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h6a.5.5 0 1 1 0 1h-6a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 0 1h-6a.5.5 0 0 1-.5-.5zm8-6a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1-.5-.5zm0 2a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1-.5-.5z"/>
                            </svg>
													</span>
                          <!--end::Svg Icon-->
												</span>
                  <span class="menu-title fw-semibold">Transactions</span>
                </router-link>
              </div>
              <!--end:Menu item-->
              <!--end:Menu sub-->
            </div>
            <!--begin:Menu item-->
            <div @click="isShow = !isShow" :class="{show: isShow}" class="menu-item menu-accordion" style="transform: unset; position: unset">
              <!--begin:Menu link-->
              <span class="menu-link">
												<span class="menu-icon">
													<!--begin::Svg Icon | path: icons/duotune/general/gen022.svg-->
													<span class="svg-icon svg-icon-2">
														<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
															<path d="M11.2929 2.70711C11.6834 2.31658 12.3166 2.31658 12.7071 2.70711L15.2929 5.29289C15.6834 5.68342 15.6834 6.31658 15.2929 6.70711L12.7071 9.29289C12.3166 9.68342 11.6834 9.68342 11.2929 9.29289L8.70711 6.70711C8.31658 6.31658 8.31658 5.68342 8.70711 5.29289L11.2929 2.70711Z" fill="currentColor"></path>
															<path d="M11.2929 14.7071C11.6834 14.3166 12.3166 14.3166 12.7071 14.7071L15.2929 17.2929C15.6834 17.6834 15.6834 18.3166 15.2929 18.7071L12.7071 21.2929C12.3166 21.6834 11.6834 21.6834 11.2929 21.2929L8.70711 18.7071C8.31658 18.3166 8.31658 17.6834 8.70711 17.2929L11.2929 14.7071Z" fill="currentColor"></path>
															<path opacity="0.3" d="M5.29289 8.70711C5.68342 8.31658 6.31658 8.31658 6.70711 8.70711L9.29289 11.2929C9.68342 11.6834 9.68342 12.3166 9.29289 12.7071L6.70711 15.2929C6.31658 15.6834 5.68342 15.6834 5.29289 15.2929L2.70711 12.7071C2.31658 12.3166 2.31658 11.6834 2.70711 11.2929L5.29289 8.70711Z" fill="currentColor"></path>
															<path opacity="0.3" d="M17.2929 8.70711C17.6834 8.31658 18.3166 8.31658 18.7071 8.70711L21.2929 11.2929C21.6834 11.6834 21.6834 12.3166 21.2929 12.7071L18.7071 15.2929C18.3166 15.6834 17.6834 15.6834 17.2929 15.2929L14.7071 12.7071C14.3166 12.3166 14.3166 11.6834 14.7071 11.2929L17.2929 8.70711Z" fill="currentColor"></path>
														</svg>
													</span>
                          <!--end::Svg Icon-->
												</span>
												<span class="menu-title">Settings</span>
												<span class="menu-arrow"></span>
											</span>
              <!--end:Menu link-->
              <!--begin:Menu sub-->
              <div @click="isShow = !isShow" :class="{show: isShow}" class="menu-sub opacity-menu menu-sub-accordion">
                <div class="opacity-menu-block">
                  <!--begin:Menu item-->
                  <div class="menu-item">
                    <!--begin:Menu link-->
                    <router-link to="/users" class="menu-link" >
														<span class="menu-bullet">
                              <span class="bullet bullet-dot"></span>
														</span>
                      <span class="menu-title">Users List</span>

                    </router-link>
                    <!--end:Menu link-->
                  </div>
                  <!--end:Menu item-->
                  <!--begin:Menu item-->
                  <div class="menu-item">
                    <!--begin:Menu link-->
                    <router-link to="/sessions" class="menu-link">
														<span class="menu-bullet">
                              <span class="bullet bullet-dot"></span>
														</span>
                      <span class="menu-title">Sessions</span>

                    </router-link>
                    <!--end:Menu link-->
                  </div>
                  <!--end:Menu item-->
                  <!--begin:Menu item-->
                  <div class="menu-item">
                    <!--begin:Menu link-->
                    <router-link to="/suppliers" class="menu-link">
														<span class="menu-bullet">
                              <span class="bullet bullet-dot"></span>
														</span>
                      <span class="menu-title">Service Providers</span>

                    </router-link>
                    <!--end:Menu link-->
                  </div>
                  <!--end:Menu item-->
                  <!--begin:Menu item-->
                  <div class="menu-item">
                    <!--begin:Menu link-->
                    <router-link to="/logs" class="menu-link">
														<span class="menu-bullet">
                              <span class="bullet bullet-dot"></span>
														</span>
                      <span class="menu-title">Action log</span>

                    </router-link>
                    <!--end:Menu link-->
                  </div>
                  <!--end:Menu item-->
                  <!--begin:Menu item-->
                  <div class="menu-item">
                    <!--begin:Menu link-->
                    <router-link to="/tasks" class="menu-link">
														<span class="menu-bullet">
                              <span class="bullet bullet-dot"></span>
														</span>
                      <span class="menu-title">Task list</span>

                    </router-link>
                    <!--end:Menu link-->
                  </div>
                  <!--end:Menu item-->
                </div>
              </div>
              <!--end:Menu sub-->
            </div>
            <!--end:Menu item-->
          </div>
          <!--end::Menu-->
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import {defineComponent} from 'vue'
import {mapGetters} from "vuex";

export default defineComponent({
  name: 'profilePage',
  data() {
    return {
      isShow: false
    }
  },

  created() {
    let path = window.location.href.split("/")[3]
    if(path === "users" || path === "suppliers" || path === "sessions") {
      this.isShow = true
    }
  },
  computed: {
    ...mapGetters({
      loggedIn: "auth/isLoggedIn",
      user: "auth/currentUser"
    }),
    currentUser() {
      const user = this.$store.getters["auth/currentUser"];
      if (user) {
        return [user.email, user.name];
      }
      return "";
    },
  },
})
</script>

<style>


.slide-fade-enter-active {
  max-height: fit-content;
  transition: max-height .1s ease;
}

.slide-fade-leave-active {
  transition: max-height .2s cubic-bezier(1.0, 1.0, 1.0, 1.0);
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateY(-20px);
  max-height: 0;
  opacity: 0;
}

.opacity-menu-block {
  opacity: 0;
  transition: all 0.2s ease !important;
}

.opacity-menu {
  max-height: 0 !important;
  transition: all 0.2s ease-out !important;
}

.show .opacity-menu-block {
  opacity: 1;
}

.show .opacity-menu {
  max-height: 100%;
}

.aside-menu .router-link-active {
  transition: color .2s ease !important;
  background-color: var(--kt-aside-menu-link-bg-color-active) !important;
}

.aside-menu .router-link-active:hover {
  color: #181C32 !important;
  background-color: var(--kt-aside-menu-link-bg-color-active) !important;
}

.aside-menu .router-link-active .menu-title {
  color: #181C32 !important;
}

.aside-menu .router-link-active .bullet-dot {
  background-color: #181C32 !important;
}
</style>