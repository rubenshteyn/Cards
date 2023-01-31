<template>
  <div v-show="isHidden" @click="isHidden = false" class="position-absolute vh-100 w-100"
       style="z-index: 103;left: 0px;top: 0px;"></div>
  <div v-show="isOpen" @click="isOpen = false" class="drawer-overlay" style="background: rgba(0, 0, 0, 0.5)"></div>
  <div id="kt_header" class="header" style="z-index: 104;">
    <!--begin::Container-->
    <div class="container-fluid d-flex flex-stack">
      <!--begin::Brand-->
      <div class="d-flex align-items-center me-5">
        <!--begin::Aside toggle-->
        <button @click="isOpen = !isOpen" class="d-lg-none btn btn-icon btn-active-color-white w-30px h-30px ms-n2 me-3"
                id="kt_aside_toggle">
          <!--begin::Svg Icon | path: icons/duotune/abstract/abs015.svg-->
          <span class="svg-icon svg-icon-2">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M21 7H3C2.4 7 2 6.6 2 6V4C2 3.4 2.4 3 3 3H21C21.6 3 22 3.4 22 4V6C22 6.6 21.6 7 21 7Z"
                    fill="currentColor"/>
              <path opacity="0.3"
                    d="M21 14H3C2.4 14 2 13.6 2 13V11C2 10.4 2.4 10 3 10H21C21.6 10 22 10.4 22 11V13C22 13.6 21.6 14 21 14ZM22 20V18C22 17.4 21.6 17 21 17H3C2.4 17 2 17.4 2 18V20C2 20.6 2.4 21 3 21H21C21.6 21 22 20.6 22 20Z"
                    fill="currentColor"/>
            </svg>
          </span>
          <!--end::Svg Icon-->
        </button>
        <!--end::Aside  toggle-->
        <!--begin::Logo-->
        <router-link @click.prevent to="/" class="p-0">
          <img alt="Logo" :src="logo" class="h-25px h-lg-30px"/>
        </router-link>
        <!--end::Logo-->
      </div>
      <!--end::Brand-->
      <!--begin::Topbar-->
      <div class="d-flex align-items-center flex-shrink-0">

        <localChanger/>

        <div class="right-links d-flex align-items-center ms-1">
          <router-link class="app-bar-item" href="#" v-if="!loggedIn" @click.prevent to="/login">Войти
          </router-link>
          <!-- <a class="app-bar-item" href="#" v-else @click.prevent="logoutButtonClicked">Выйти</a> -->
          <!--begin::User-->
          <div @click='isHidden = !isHidden' v-else class="d-flex align-items-center ms-1">
            <!--begin::User info-->
            <div class="btn btn-flex align-items-center bg-hover-white bg-hover-opacity-10 py-2 px-2 px-md-3"
                 data-kt-menu-trigger="click" data-kt-menu-attach="parent" data-kt-menu-placement="bottom-end">
              <!--begin::Name-->
              <div class="d-none d-md-flex flex-column align-items-end justify-content-center me-2 me-md-4">
                <span class="text-muted fs-8 fw-semibold lh-1 mb-1">{{ currentUser()[1] }}</span>
              </div>
              <!--end::Name-->
              <!--begin::Symbol-->
              <div class="symbol symbol-30px symbol-md-40px">
                <img :src="avatar"/>
              </div>
              <!--end::Symbol-->
            </div>
            <!--end::User info-->
            <!--begin::User account menu-->
            <div v-show="isHidden"
                 class="menu user-profile-menu show menu-sub menu-sub-dropdown menu-column menu-rounded menu-gray-800 menu-state-bg menu-state-color fw-semibold py-4 fs-6 w-275px"
                 style="z-index: 999; right: 0 !important;">
              <!--begin::Menu item-->
              <div class="menu-item px-3">
                <div class="menu-content d-flex align-items-center px-3">
                  <!--begin::Avatar-->
                  <div class="symbol symbol-50px me-5">
                    <img :src="avatar"/>
                  </div>
                  <!--end::Avatar-->
                  <!--begin::Username-->
                  <div class="d-flex flex-column">
                    <div class="fw-bold d-flex align-items-center fs-5">{{ currentUser()[1] }}
                      <span v-if="currentUser()[2] === 2" class="badge badge-light-success fw-bold fs-8 px-2 py-1 ms-2">Admin</span>
                      <span v-if="currentUser()[2] === 1" class="badge badge-light-success fw-bold fs-8 px-2 py-1 ms-2">User</span>
                      <span v-if="currentUser()[2] === 3" class="badge badge-light-primary fw-bold fs-8 px-2 py-1 ms-2">Manager</span>
                    </div>
                    <span class="fw-semibold text-muted text-hover-primary fs-7">{{ currentUser()[0] }}</span>
                  </div>
                  <!--end::Username-->
                </div>
              </div>
              <!--end::Menu item-->
              <!--begin::Menu separator-->
              <div class="separator my-2"></div>
              <!--end::Menu separator-->
              <!--begin::Menu item-->
              <div class="menu-item px-5">
                <router-link to="/profile" class="menu-link px-5">{{ $t('header.Profile') }}</router-link>
              </div>
              <!--end::Menu item-->
              <!--begin::Menu item-->
<!--              <div class="menu-item px-5">-->
<!--                <span class="menu-link px-5">{{ $t('header.settings') }}</span>-->
<!--              </div>-->
              <!--end::Menu item-->
              <!--begin::Menu item-->
              <div class="menu-item px-5">
                <router-link to="/sessions" class="menu-link px-5"
                             style="margin-left: auto">{{ $t('sessions.tableTitle') }}</router-link>
              </div>
              <!--end::Menu item-->
              <!--begin::Menu separator-->
              <div class="separator my-2"></div>
              <!--end::Menu separator-->

              <!--begin::Menu item-->
              <div class="menu-item px-5">
                <span @click.prevent="logoutButtonClicked" class="menu-link px-5"
                      style="margin-left: auto">{{ $t('header.logout') }}</span>
              </div>
              <!--end::Menu item-->
            </div>
            <!--end::User account menu-->
          </div>
        </div>
        <!--end::User -->
      </div>
      <!--end::Topbar-->
    </div>
    <!--end::Container-->
  </div>
  <div id="kt_aside" class="aside card drawer drawer-start drawer-on" :style="{transform: showMenu()}" data-kt-drawer="true" data-kt-drawer-name="aside"
       data-kt-drawer-activate="{default: true, lg: false}" data-kt-drawer-overlay="true"
       data-kt-drawer-width="{default:'200px', '300px': '250px'}" data-kt-drawer-direction="start"
       data-kt-drawer-toggle="#kt_aside_toggle" style="width: 250px !important;">
    <!--begin::Aside menu-->
    <div class="aside-menu flex-column-fluid px-4">
      <!--begin::Aside Menu-->
      <div class="hover-scroll-overlay-y my-5 pe-4 me-n4" id="kt_aside_menu_wrapper" data-kt-scroll="true"
           data-kt-scroll-activate="true" data-kt-scroll-height="auto"
           data-kt-scroll-dependencies="{default: '#kt_aside_footer', lg: '#kt_header, #kt_aside_footer'}"
           data-kt-scroll-wrappers="#kt_aside, #kt_aside_menu" data-kt-scroll-offset="{default: '5px', lg: '75px'}"
           style="height: 616px;">
        <!--begin::Menu-->
        <div class="menu menu-column menu-rounded menu-sub-indention fw-semibold fs-6" id="#kt_aside_menu"
             data-kt-menu="true">
          <!--begin:Menu item-->
          <div data-kt-menu-trigger="click" class="menu-item menu-accordion show">
            <!--begin:Menu sub-->
            <!--begin:Menu item-->
<!--            <div class="menu-item">-->
<!--              <router-link to="/dashboard" class="menu-link">-->
<!--												<span class="menu-icon">-->
<!--													&lt;!&ndash;begin::Svg Icon | path: icons/duotune/general/gen025.svg&ndash;&gt;-->
<!--													<span class="svg-icon svg-icon-2 text-gray-400">-->
<!--														<svg width="24" height="24" viewBox="0 0 24 24" fill="none"-->
<!--                                 xmlns="http://www.w3.org/2000/svg">-->
<!--															<rect x="2" y="2" width="9" height="9" rx="2" fill="currentColor"></rect>-->
<!--															<rect opacity="0.3" x="13" y="2" width="9" height="9" rx="2" fill="currentColor"></rect>-->
<!--															<rect opacity="0.3" x="13" y="13" width="9" height="9" rx="2" fill="currentColor"></rect>-->
<!--															<rect opacity="0.3" x="2" y="13" width="9" height="9" rx="2" fill="currentColor"></rect>-->
<!--														</svg>-->
<!--													</span>-->
<!--                          &lt;!&ndash;end::Svg Icon&ndash;&gt;-->
<!--												</span>-->
<!--                <span class="menu-title fw-semibold">Dashboard</span>-->
<!--              </router-link>-->
<!--            </div>-->
            <!--end:Menu item-->
<!--            &lt;!&ndash;begin:Menu item&ndash;&gt;-->
<!--            <div class="menu-item">-->
<!--              <router-link @click="isOpen = !isOpen" to="/accounts" class="menu-link">-->
<!--												<span class="menu-icon">-->
<!--													&lt;!&ndash;begin::Svg Icon | path: icons/duotune/general/gen022.svg&ndash;&gt;-->
<!--													<span class="svg-icon svg-icon-2">-->
<!--														<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">-->
<!--															<path d="M11.2929 2.70711C11.6834 2.31658 12.3166 2.31658 12.7071 2.70711L15.2929 5.29289C15.6834 5.68342 15.6834 6.31658 15.2929 6.70711L12.7071 9.29289C12.3166 9.68342 11.6834 9.68342 11.2929 9.29289L8.70711 6.70711C8.31658 6.31658 8.31658 5.68342 8.70711 5.29289L11.2929 2.70711Z" fill="currentColor"></path>-->
<!--															<path d="M11.2929 14.7071C11.6834 14.3166 12.3166 14.3166 12.7071 14.7071L15.2929 17.2929C15.6834 17.6834 15.6834 18.3166 15.2929 18.7071L12.7071 21.2929C12.3166 21.6834 11.6834 21.6834 11.2929 21.2929L8.70711 18.7071C8.31658 18.3166 8.31658 17.6834 8.70711 17.2929L11.2929 14.7071Z" fill="currentColor"></path>-->
<!--															<path opacity="0.3" d="M5.29289 8.70711C5.68342 8.31658 6.31658 8.31658 6.70711 8.70711L9.29289 11.2929C9.68342 11.6834 9.68342 12.3166 9.29289 12.7071L6.70711 15.2929C6.31658 15.6834 5.68342 15.6834 5.29289 15.2929L2.70711 12.7071C2.31658 12.3166 2.31658 11.6834 2.70711 11.2929L5.29289 8.70711Z" fill="currentColor"></path>-->
<!--															<path opacity="0.3" d="M17.2929 8.70711C17.6834 8.31658 18.3166 8.31658 18.7071 8.70711L21.2929 11.2929C21.6834 11.6834 21.6834 12.3166 21.2929 12.7071L18.7071 15.2929C18.3166 15.6834 17.6834 15.6834 17.2929 15.2929L14.7071 12.7071C14.3166 12.3166 14.3166 11.6834 14.7071 11.2929L17.2929 8.70711Z" fill="currentColor"></path>-->
<!--														</svg>-->
<!--													</span>-->
<!--                          &lt;!&ndash;end::Svg Icon&ndash;&gt;-->
<!--												</span>-->
<!--                <span class="menu-title">{{ $t('menu.account') }}</span>-->
<!--              </router-link>-->
<!--            </div>-->
<!--            &lt;!&ndash;end:Menu item&ndash;&gt;-->
            <!--begin:Menu item-->
            <div class="menu-item">
              <router-link @click="isOpen = !isOpen" to="/cards" class="menu-link">
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
                <span class="menu-title fw-semibold">{{ $t('menu.cardList') }}</span>
              </router-link>
            </div>
            <!--end:Menu item-->
            <!--begin:Menu item-->
            <div class="menu-item">
              <router-link @click="isOpen = !isOpen" to="/transactions" class="menu-link">
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
                <span class="menu-title fw-semibold">{{ $t('menu.transactions') }}</span>
              </router-link>
            </div>
            <!--end:Menu item-->
            <!--begin:Menu item-->
            <div class="menu-item">
              <router-link @click="isOpen = !isOpen" to="/tasks" class="menu-link">
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
                <span class="menu-title fw-semibold">{{ $t('menu.tasks') }}</span>
              </router-link>
            </div>
            <!--end:Menu item-->
            <!--end:Menu sub-->
          </div>
        </div>
        <!--end::Menu-->
      </div>
    </div>
    <!--end::Aside menu-->
  </div>
</template>

<script>
import {mapActions, mapGetters} from "vuex";
import logo from '@/assets/media/logo.png'
import avatar from '@/assets/media/001-boy.svg'
import localChanger from "@/components/localeChanger";

export default {
  name: "mainHeader",
  data() {
    return {
      logo,
      avatar,
      isShow: false,
      isOpen: false,
      isHidden: false,
    }
  },
  methods: {
    ...mapActions({
      logout: "auth/logout"
    }),
    currentUser() {
      const user = this.$store.getters["auth/currentUser"];
      if (user) {
        return [user.email, user.name, user.role];
      }
      return "";
    },
    showMenu() {
      // if (window.innerWidth < 990) {
      //   let x = null;
      //   document.addEventListener('touchstart', e => x = e.touches[0].clientX);
      //   document.addEventListener('touchmove', e => {
      //     if (!x) return
      //     x = x - e.touches[0].clientX < 0 ? 0 : -250
      //     document.querySelector(".drawer").style.transform = `translate(${x}%,0)`
      //     if (x !== 0) {
      //       this.isOpen = false
      //     }
      //     if (x === 0) {
      //       this.isOpen = true
      //     }
      //     x = null
      //   });
      // }
      if (this.isOpen === true) {
        return "none"
      }
      return "translateX(-250%)"
    },
    logoutButtonClicked() {
      this.logout();
    },
  },
  computed: {
    ...mapGetters({
      loggedIn: "auth/isLoggedIn",
      user: "auth/currentUser"
    }),
    links() {
      return [{
        visibleIfLoggedOut: false,
        name: "Основная страница",
        to: "/"
      },
        {
          visibleIfLoggedOut: false,
          name: "Профиль",
          to: "/profile"
        }
      ];
    },
    activeLinks() {
      return this.links.filter(
          link => link.visibleIfLoggedOut || this.loggedIn
      );
    }
  },
  components: {
    localChanger
  }
};
</script>

<style scoped>
.aside-menu .router-link-active {
  transition: color .2s ease;
  background-color: var(--kt-aside-menu-link-bg-color-active);
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

.aside-menu .router-link-active .menu-title {
  color: #181C32 !important;
}
.drawer .show {
  z-index: unset;
  position: unset;
  inset: inherit;
  margin: unset;
  transform: unset;
}
.show {
  z-index: 105;
  position: fixed;
  inset: 0px 0px auto auto;
  margin: 0px;
  transform: translate(-50px, 62px);
}

.app-bar {
  height: 5vh;
  width: 100%;
  position: fixed;
  z-index: 1;
  top: 0;
  left: 0;
  background-color: burlywood;
  overflow: hidden;
  display: block;
  justify-content: space-between;
}

.left-links {
  padding-left: 10%;
}

.left-links a {
  float: left;
}

.right-links {
  float: right;
}
</style>