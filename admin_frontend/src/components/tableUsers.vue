<template>
  <div v-show="isHidden" @click="isHidden = false" class="position-absolute vh-100 w-100"
       style="z-index: 99;left: 0px;top: 0px;"></div>
  <div v-show="actionsShow" @click="actionsShow = false" class="position-absolute vh-100 w-100"
       style="z-index: 99;left: 0px;top: 0px;"></div>
  <div class="d-flex flex-column flex-column-fluid container-fluid">
    <div class="toolbar mb-5 mb-lg-7" id="kt_toolbar">
      <!--begin::Page title-->
      <div class="page-title d-flex flex-column me-3">
        <!--begin::Title-->
        <h1 class="d-flex text-dark fw-bold my-1 fs-3">Users management</h1>
        <!--end::Title-->
        <!--begin::Breadcrumb-->
        <ul class="breadcrumb breadcrumb-dot fw-semibold text-gray-600 fs-7 my-1">
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-600">
            <router-link to="/" class="text-gray-600 text-hover-primary">Home</router-link>
          </li>
          <!--end::Item-->
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-500">Users management</li>
          <!--end::Item-->
        </ul>
        <!--end::Breadcrumb-->
      </div>
      <!--end::Page title-->
      <!--begin::Actions-->
      <div class="actions-group-buttons d-flex align-items-center py-2 py-md-1 position-relative">
        <!--begin::Wrapper-->
        <div class="me-3">
          <!--begin::Menu-->
          <button @click='isHidden = !isHidden' class="btn btn-light fw-bold">
            <!--begin::Svg Icon | path: icons/duotune/general/gen031.svg-->
            <span class="svg-icon svg-icon-5 svg-icon-gray-500 me-1">
                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                 xmlns="http://www.w3.org/2000/svg">
                                <path
                                    d="M19.0759 3H4.72777C3.95892 3 3.47768 3.83148 3.86067 4.49814L8.56967 12.6949C9.17923 13.7559 9.5 14.9582 9.5 16.1819V19.5072C9.5 20.2189 10.2223 20.7028 10.8805 20.432L13.8805 19.1977C14.2553 19.0435 14.5 18.6783 14.5 18.273V13.8372C14.5 12.8089 14.8171 11.8056 15.408 10.964L19.8943 4.57465C20.3596 3.912 19.8856 3 19.0759 3Z"
                                    fill="currentColor"/>
                            </svg>
                        </span>
            <!--end::Svg Icon-->
            <span class="user-filter-span">Filter</span>
          </button>
          <!--begin::Menu 1-->
          <div id="filter-modal" v-show="isHidden"
               class="menu menu-sub menu-sub-dropdown show position-fixed float-left" data-kt-menu="true"
               style="width: inherit; z-index: 999; right: 50px;">
            <!--begin::Header-->
            <div class="px-7 py-5">
              <div class="fs-5 text-dark fw-bold">Users filter</div>
            </div>
            <!--end::Header-->
            <!--begin::Menu separator-->
            <div class="separator border-gray-200"></div>
            <!--end::Menu separator-->
            <!--begin::Form-->
            <div class="px-7 py-5">
              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="status">User status:</label>
                  <v-select class="bg-transparent" label="name" :options="filterItems.statuses"
                            v-model="filterItems.status" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="role">User role:</label>
                  <v-select class="bg-transparent" label="name" :options="filterItems.roles"
                            v-model="filterItems.role" required>
                  </v-select>
                </div>
              </div>

              <div class="mb-10">
                <div class="form-group fv-row">
                  <label class="form-label fw-semibold" for="search">User search</label>
                  <input class="form-control-solid form-control form-control-lg" label="name"
                         v-model="filterItems.search">
                </div>
              </div>


              <div class="d-flex justify-content-end">
                <button @click="resetFilterUsers" type="reset"
                        class="btn btn-sm btn-light btn-active-light-primary me-2">Reset
                </button>
                <button @click="filter" type="submit" class="btn btn-sm btn-primary"
                        data-kt-menu-dismiss="true">Apply
                </button>
              </div>
              <!--end::Actions-->
            </div>
            <!--end::Form-->
          </div>
          <!--end::Menu 1-->
          <!--end::Menu-->
        </div>
        <!--end::Wrapper-->
        <Popper class="fw-semibold fs-7" arrow hover placement="top" content="Click to create user">
          <button @click='showModal' class="btn btn-dark fw-bold" data-bs-toggle="modal"
                  data-bs-target="#kt_modal_create_app" id="kt_toolbar_primary_button">Create user
          </button>
        </Popper>
        <CreateUserModal ref="modal"/>
      </div>
      <!--end::Actions-->
    </div>
    <div class="card mb-5 mb-xl-8">
      <!--begin::Header-->
      <div class="card-header align-items-center border-0 pt-5">
        <h3 class="card-title align-items-start flex-column">
          <span class="card-label fw-bold fs-3 mb-1">Users</span>
          <span class="text-muted mt-1 fw-semibold fs-7">Amount users: {{ this.filteredUsers.length }}</span>
        </h3>

<!--        <button @click="actionsShow = !actionsShow" type="button"-->
<!--                class="d-none btn open-actions-card-btn btn-sm btn-icon btn-color-primary btn-active-light-primary"-->
<!--                data-kt-menu-trigger="click" data-kt-menu-placement="bottom-end">-->
<!--          &lt;!&ndash;begin::Svg Icon | path: icons/duotune/general/gen024.svg&ndash;&gt;-->
<!--          <span class="svg-icon svg-icon-2">-->
<!--                            <svg xmlns="http://www.w3.org/2000/svg" width="24px" height="24px" viewBox="0 0 24 24">-->
<!--                                <g stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">-->
<!--                                    <rect x="5" y="5" width="5" height="5" rx="1" fill="currentColor"/>-->
<!--                                    <rect x="14" y="5" width="5" height="5" rx="1" fill="currentColor" opacity="0.3"/>-->
<!--                                    <rect x="5" y="14" width="5" height="5" rx="1" fill="currentColor" opacity="0.3"/>-->
<!--                                    <rect x="14" y="14" width="5" height="5" rx="1" fill="currentColor" opacity="0.3"/>-->
<!--                                </g>-->
<!--                            </svg>-->
<!--                        </span>-->
<!--          &lt;!&ndash;end::Svg Icon&ndash;&gt;-->
<!--        </button>-->

<!--        <div v-if="actionsShow" class="card-toolbar card-toolbar-card">-->
<!--          <button @click="showBalanceModal()"-->
<!--                  class="btn show-delete btn-icon btn-bg-light btn-active-color-primary btn-sm">-->
<!--            &lt;!&ndash;begin::Svg Icon | path: C:/wamp64/www/keenthemes/core/html/src/media/icons/duotune/finance/fin008.svg&ndash;&gt;-->
<!--            <span class="svg-icon svg-icon-muted svg-icon-2hx" style="margin-right: 0">-->
<!--                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">-->
<!--                <path opacity="0.3" d="M3.20001 5.91897L16.9 3.01895C17.4 2.91895 18 3.219 18.1 3.819L19.2 9.01895L3.20001 5.91897Z" fill="currentColor"/>-->
<!--                <path opacity="0.3" d="M13 13.9189C13 12.2189 14.3 10.9189 16 10.9189H21C21.6 10.9189 22 11.3189 22 11.9189V15.9189C22 16.5189 21.6 16.9189 21 16.9189H16C14.3 16.9189 13 15.6189 13 13.9189ZM16 12.4189C15.2 12.4189 14.5 13.1189 14.5 13.9189C14.5 14.7189 15.2 15.4189 16 15.4189C16.8 15.4189 17.5 14.7189 17.5 13.9189C17.5 13.1189 16.8 12.4189 16 12.4189Z" fill="currentColor"/>-->
<!--                <path d="M13 13.9189C13 12.2189 14.3 10.9189 16 10.9189H21V7.91895C21 6.81895 20.1 5.91895 19 5.91895H3C2.4 5.91895 2 6.31895 2 6.91895V20.9189C2 21.5189 2.4 21.9189 3 21.9189H19C20.1 21.9189 21 21.0189 21 19.9189V16.9189H16C14.3 16.9189 13 15.6189 13 13.9189Z" fill="currentColor"/>-->
<!--                </svg>-->
<!--              </span>-->
<!--            &lt;!&ndash;end::Svg Icon&ndash;&gt;-->
<!--          </button>-->
<!--        </div>-->

        <div class="card-toolbar card-toolbar-card-desktop">
          <Popper class="fw-semibold fs-7" arrow hover placement="top" content="Click to top up balance">
            <button @click="showBalanceModal()"
                    class="btn show-delete btn-icon btn-bg-light btn-active-color-primary btn-sm">
              <!--begin::Svg Icon | path: C:/wamp64/www/keenthemes/core/html/src/media/icons/duotune/finance/fin008.svg-->
              <span class="svg-icon svg-icon-muted svg-icon-2hx" style="margin-right: 0">
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path opacity="0.3" d="M3.20001 5.91897L16.9 3.01895C17.4 2.91895 18 3.219 18.1 3.819L19.2 9.01895L3.20001 5.91897Z" fill="currentColor"/>
                <path opacity="0.3" d="M13 13.9189C13 12.2189 14.3 10.9189 16 10.9189H21C21.6 10.9189 22 11.3189 22 11.9189V15.9189C22 16.5189 21.6 16.9189 21 16.9189H16C14.3 16.9189 13 15.6189 13 13.9189ZM16 12.4189C15.2 12.4189 14.5 13.1189 14.5 13.9189C14.5 14.7189 15.2 15.4189 16 15.4189C16.8 15.4189 17.5 14.7189 17.5 13.9189C17.5 13.1189 16.8 12.4189 16 12.4189Z" fill="currentColor"/>
                <path d="M13 13.9189C13 12.2189 14.3 10.9189 16 10.9189H21V7.91895C21 6.81895 20.1 5.91895 19 5.91895H3C2.4 5.91895 2 6.31895 2 6.91895V20.9189C2 21.5189 2.4 21.9189 3 21.9189H19C20.1 21.9189 21 21.0189 21 19.9189V16.9189H16C14.3 16.9189 13 15.6189 13 13.9189Z" fill="currentColor"/>
                </svg>
              </span>
              <!--end::Svg Icon-->
            </button>
          </Popper>
        </div>
      </div>
      <!--end::Header-->
      <!--begin::Body-->
      <div class="card-body py-3">
        <!--begin::Table container-->
        <div class="table-responsive">
          <!--begin::Table-->
          <table v-if="getUsers()" class="table table-row-bordered table-row-gray-100 align-middle gs-0 gy-3">
            <!--begin::Table head-->
            <thead class="rounded-left">
            <tr class="fw-bold text-muted bg-light">
              <th class="w-25px">
                <div class="form-check form-check-sm form-check-custom form-check-solid mr-3">
                  <input @click="selectAllUsers" v-model="check" class="form-check-input" type="checkbox"/>
                </div>
              </th>
              <th class="min-w-170px">User</th>
              <th @click="filterUsersBalance('USDTBalance')" class="min-w-120px cursor-pointer text-hover-primary">
                USDT
              </th>
              <th @click="filterUsersBalance('BTCBalance')" class="min-w-120px cursor-pointer text-hover-primary">BTC
              </th>
              <th @click="filterUsersBalance('USDBalance')" class="min-w-120px cursor-pointer text-hover-primary">
                Dollar
              </th>
              <th @click="filterUsersBalance('EURBalance')" class="min-w-120px cursor-pointer text-hover-primary">Euro
              </th>
              <th class="min-w-150px">Country</th>
              <th class="min-w-160px">Contact</th>
              <th @click="filterUserDate('CreatedAt')" class="min-w-150px cursor-pointer text-hover-primary">
                Registration
              </th>
              <th class="min-w-120px">Status</th>
              <th class="min-w-120px">Role</th>
              <th class="min-w-100px text-end">Actions</th>
            </tr>
            </thead>
            <!--end::Table head-->
            <!--begin::Table body-->
            <tbody>
            <tr v-for="user in filteredUsers" :key="user" :style="{background: background}">
              <td class="align-middle">
                <div class="form-check form-check-sm form-check-custom form-check-solid">
                  <input v-bind:value="user" v-model="checkedUsers" class="form-check-input widget-9-check"
                         type="checkbox"/>
                </div>
              </td>
              <td class="align-middle">
                <div class="d-flex align-items-center">
                  <div class="symbol symbol-50px me-5">
                                            <span class="symbol-label bg-light">
                                                <img src="../assets/media/001-boy.svg" class="h-75 align-self-end"
                                                     alt=""/>
                                            </span>
                  </div>
                  <div class="d-flex justify-content-start flex-column">
                    <span @click="showEditModal(user)"
                          class="text-dark fw-bold text-hover-primary fs-6 cursor-pointer text-hover-primary mb-1" style="white-space: nowrap">{{
                        user.name
                      }}, ID: {{user.ID}}</span>
                    <a v-bind:href="'mailto:' + user.email "
                       class="text-muted fw-semibold text-muted d-block fs-7">{{ user.email }}</a>
                  </div>
                </div>
              </td>
              <td @click="showBalanceModalUser(user, user.USDTBalance, 'USDT', '₮')" class="cursor-pointer align-middle">
                <div class="d-flex flex-column">
                  <span class="text-dark fw-bold d-block fs-6 mb-1">
                  {{ userBalance(user.USDTBalance) + "₮" }}
                </span>
                  <span class="text-muted fw-semibold text-muted d-block fs-7">
                    ~XXX$
                </span>
                </div>
              </td>
              <td @click="showBalanceModalUser(user, user.BTCBalance, 'BTC', '₿')" class="cursor-pointer align-middle">
                <div class="d-flex flex-column">
                  <span class="text-dark fw-bold d-block fs-6 mb-1">
                  {{ userBalance(user.BTCBalance) + "₿" }}
                </span>
                  <span class="text-muted fw-semibold text-muted d-block fs-7">
                    ~XXX$
                </span>
                </div>
              </td>
              <td @click="showBalanceModalUser(user, user.USDBalance, 'USD', '$')" class="cursor-pointer align-middle">
                <div class="d-flex flex-column">
                  <span class="text-dark fw-bold d-block fs-6 mb-1">
                  {{ userBalance(user.USDBalance) + "$" }}
                </span>
                  <span class="text-muted fw-semibold text-muted d-block fs-7">
                    ~XXX$
                </span>
                </div>
              </td>
              <td @click="showBalanceModalUser(user, user.EURBalance, 'EUR', '€')" class="cursor-pointer align-middle">
                <div class="d-flex flex-column">
                  <span class="text-dark fw-bold d-block fs-6 mb-1">
                  {{ userBalance(user.EURBalance) + "€" }}
                </span>
                  <span class="text-muted fw-semibold text-muted d-block fs-7">
                    ~XXX$
                </span>
                </div>
              </td>
              <td class="align-middle">
                <div class="d-flex flex-column">
                  <span class="text-dark fw-bold d-block fs-6 mb-1">{{ getCountry(user.country) }}</span>
                  <span
                      class="text-muted fw-semibold text-muted d-block fs-7">{{ user.lastIP }}</span>

                </div>
              </td>

              <td class="align-middle">
                <div>
                  <a class="copy-btn d-flex align-items-center cursor-pointer"
                     @click="copySkype(includeMessenger(user))" style="color: #009ef7">
                    <svg width="23" height="23" viewBox="0 0 23 23" fill="none" xmlns="http://www.w3.org/2000/svg"
                         v-if="includeMessenger(user)[1] === 'TG'">
                      <path
                          d="M11.5 23C17.8513 23 23 17.8513 23 11.5C23 5.14873 17.8513 0 11.5 0C5.14873 0 0 5.14873 0 11.5C0 17.8513 5.14873 23 11.5 23Z"
                          fill="#039BE5"/>
                      <path
                          d="M5.26234 11.2509L16.3503 6.97578C16.8649 6.78987 17.3143 7.10133 17.1476 7.87949L17.1486 7.87854L15.2606 16.7728C15.1207 17.4034 14.746 17.5567 14.2218 17.2597L11.3468 15.1408L9.96009 16.4767C9.80676 16.63 9.67738 16.7594 9.3803 16.7594L9.58442 13.8336L14.9128 9.01991C15.1447 8.81579 14.861 8.70079 14.5553 8.90395L7.97059 13.0497L5.132 12.1642C4.5158 11.9687 4.50238 11.548 5.26234 11.2509Z"
                          fill="white"/>
                    </svg>
                    <svg v-if="includeMessenger(user)[1] === 'SKYPE'" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="25px" height="25px">
                      <path fill="#03a9f4" d="M24 4A20 20 0 1 0 24 44A20 20 0 1 0 24 4Z"/>
                      <path fill="#03a9f4"
                            d="M33.5 22A11.5 11.5 0 1 0 33.5 45 11.5 11.5 0 1 0 33.5 22zM14.5 3A11.5 11.5 0 1 0 14.5 26 11.5 11.5 0 1 0 14.5 3z"/>
                      <path fill="#fff"
                            d="M24.602,36C18,36,15,32.699,15,30.199C15,28.898,15.898,28,17.199,28c2.801,0,2.102,4.102,7.402,4.102c2.699,0,4.199-1.5,4.199-3c0-0.902-0.402-1.902-2.199-2.402l-5.902-1.5C16,24,15.102,21.398,15.102,18.898c0-5.098,4.699-6.898,9.098-6.898C28.301,12,33,14.199,33,17.199c0,1.301-1,2.102-2.301,2.102c-2.398,0-2-3.402-6.801-3.402c-2.398,0-3.797,1.102-3.797,2.703c0,1.598,1.898,2.098,3.598,2.5l4.402,1C32.898,23.199,34,26,34,28.699C33.898,32.898,30.898,36,24.602,36z"/>
                    </svg>
                    <span class="text-dark" v-if="includeMessenger(user)[1] === ''">-</span>
                    <span class="copy-text d-block fw-semibold ml-1 fs-7">{{ includeMessenger(user)[2] }}</span>
                  </a>
                </div>
              </td>
              <td class="text-muted fw-semibold fs-6 align-middle">
                <div class="d-flex flex-column">
                  <span class="text-muted fw-semibold text-muted d-block fs-7 mb-1">
                    {{ splitDate(user.CreatedAt) }}
                  </span>
                  <span class="text-muted fw-semibold text-muted d-block fs-7">
                    {{ splitDate(user.lastActionDate) }}
                  </span>
                </div>
              </td>
              <td class="align-middle">
                                    <span v-if="user.status === 0" class="badge badge-light-danger">
                                        Unverified email
                                    </span>

                <span v-if="user.status === 1" class="badge badge-light-warning">
                                        Moderation
                                    </span>

                <span v-if="user.status === 2" class="badge badge-light-success">
                                        Active
                                    </span>
                <span v-if="user.status === 3" class="badge badge-light-danger">
                                        Blocked
                                    </span>
              </td>
              <td class="align-middle">
                                    <span v-if="user.role === 0" class="badge badge-light-success">
                                        Invited
                                    </span>

                <span v-if="user.role === 1" class="badge badge-light-success">
                                        User
                                    </span>

                <span v-if="user.role === 2" class="badge badge-light-success">
                                        Administrator
                                    </span>

                <span v-if="user.role === 3" class="badge badge-light-primary">
                                        Manager
                                    </span>
              </td>
              <td class="border-0 end-buttons align-middle">
                <div class="d-flex justify-content-end flex-shrink-0" style="margin-left: auto">
                  <div class="me-2">
                    <Popper class="fw-semibold fs-7" arrow hover placement="top" content="Click to login as this user">
                      <button @click="access(user)" class="btn btn-icon btn-bg-light btn-active-color-primary btn-sm">
                        <!--begin::Svg Icon | path: C:/wamp64/www/keenthemes/core/html/src/media/icons/duotune/arrows/arr076.svg-->
                        <span class="svg-icon svg-icon-muted svg-icon-3"><svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                          <rect opacity="0.3" width="12" height="2" rx="1" transform="matrix(-1 0 0 1 15.5 11)" fill="currentColor"/>
                          <path d="M13.6313 11.6927L11.8756 10.2297C11.4054 9.83785 11.3732 9.12683 11.806 8.69401C12.1957 8.3043 12.8216 8.28591 13.2336 8.65206L16.1592 11.2526C16.6067 11.6504 16.6067 12.3496 16.1592 12.7474L13.2336 15.3479C12.8216 15.7141 12.1957 15.6957 11.806 15.306C11.3732 14.8732 11.4054 14.1621 11.8756 13.7703L13.6313 12.3073C13.8232 12.1474 13.8232 11.8526 13.6313 11.6927Z" fill="currentColor"/>
                          <path d="M8 5V6C8 6.55228 8.44772 7 9 7C9.55228 7 10 6.55228 10 6C10 5.44772 10.4477 5 11 5H18C18.5523 5 19 5.44772 19 6V18C19 18.5523 18.5523 19 18 19H11C10.4477 19 10 18.5523 10 18C10 17.4477 9.55228 17 9 17C8.44772 17 8 17.4477 8 18V19C8 20.1046 8.89543 21 10 21H19C20.1046 21 21 20.1046 21 19V5C21 3.89543 20.1046 3 19 3H10C8.89543 3 8 3.89543 8 5Z" fill="currentColor"/>
                        </svg>
                        </span>
                        <!--end::Svg Icon-->
                      </button>
                    </Popper>
                  </div>
                  <Popper class="fw-semibold fs-7" arrow hover placement="top" content="Click to change user">
                    <button @click="showEditModal(user)"
                            class="btn btn-icon btn-bg-light btn-active-color-primary btn-sm">
                      <!--begin::Svg Icon | path: icons/duotune/art/art005.svg-->
                      <span class="svg-icon svg-icon-3">
																		<svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                         xmlns="http://www.w3.org/2000/svg">
																			<path opacity="0.3"
                                            d="M21.4 8.35303L19.241 10.511L13.485 4.755L15.643 2.59595C16.0248 2.21423 16.5426 1.99988 17.0825 1.99988C17.6224 1.99988 18.1402 2.21423 18.522 2.59595L21.4 5.474C21.7817 5.85581 21.9962 6.37355 21.9962 6.91345C21.9962 7.45335 21.7817 7.97122 21.4 8.35303ZM3.68699 21.932L9.88699 19.865L4.13099 14.109L2.06399 20.309C1.98815 20.5354 1.97703 20.7787 2.03189 21.0111C2.08674 21.2436 2.2054 21.4561 2.37449 21.6248C2.54359 21.7934 2.75641 21.9115 2.989 21.9658C3.22158 22.0201 3.4647 22.0084 3.69099 21.932H3.68699Z"
                                            fill="currentColor"></path>
																			<path
                                          d="M5.574 21.3L3.692 21.928C3.46591 22.0032 3.22334 22.0141 2.99144 21.9594C2.75954 21.9046 2.54744 21.7864 2.3789 21.6179C2.21036 21.4495 2.09202 21.2375 2.03711 21.0056C1.9822 20.7737 1.99289 20.5312 2.06799 20.3051L2.696 18.422L5.574 21.3ZM4.13499 14.105L9.891 19.861L19.245 10.507L13.489 4.75098L4.13499 14.105Z"
                                          fill="currentColor"></path>
																		</svg>
																	</span>
                      <!--end::Svg Icon-->
                    </button>
                  </Popper>
                </div>
              </td>
            </tr>
            </tbody>
            <!--end::Table body-->
          </table>
          <EditUserModal ref="edit"/>
          <EditBalanceModal ref="balance"/>
          <!--end::Table-->
        </div>
        <!--end::Table container-->
      </div>
      <!--begin::Body-->
    </div>
  </div>
</template>

<script>
import {mapGetters, mapActions} from 'vuex'
import countries from '../countries.json'
import timezones from '../timezones.json'
import CreateUserModal from '@/components/modals/createUserModal.vue'
import EditUserModal from './modals/editUserModal.vue';
import EditBalanceModal from './modals/editBalanceModal.vue';
import apiV1 from "@/axios";

export default {
  name: 'tableUsers',
  data() {
    return {
      timeZone: {
        id: null,
        text: null,
        utc: 0
      },
      timezones: timezones,
      background: "none",
      onTheTable: false,
      check: false,
      actionsShow: false,
      checkedUsers: [],
      isHidden: false,
      copyNick: false,
      req: true,
      filteredUsers: this.users,
      filterItems: {
        CreatedAt: 1,
        USDTBalance: 0,
        BTCBalance: 0,
        USDBalance: 0,
        EURBalance: 0,
        countries: countries,
        statuses: [
          {
            name: "Active",
            code: 2
          },
          {
            name: "Moderation",
            code: 1
          },
          {
            name: "Unverified email",
            code: 0
          },
          {
            name: "Blocked",
            code: 3
          },
        ],
        status: {
          name: null,
          code: null
        },
        roles: [
          {
            name: "Administrator",
            code: 2
          },
          {
            name: "User",
            code: 1
          },
          {
            name: "Manager",
            code: 3
          },
          {
            name: "Invited",
            code: 0
          }
        ],
        role: {
          name: null,
          code: null
        },
        country: {
          name: null,
          code: null
        },
        search: '',
      },
    }
  },
  created() {
    this.filteredUsers = this.users
    if (this.users) {
      this.filter()
    }
  },
  watch: {
    users: "filter"
  },
  methods: {
    ...mapActions({
      deleteUserMethod: "admin/deleteUser",
      login: "auth/login",
      getAccess: "admin/getAccess"
    }),
    showDeleteSelectUsers() {
      if (this.checkedUsers.length > 0) {
        return false
      }
      return true
    },
    filterUserDate(date) {
      this.filterItems[date] !== 0 ? this.filterItems[date] = -this.filterItems[date] : this.filterItems[date] = 1
      date === "CreatedAt" ? this.filterItems.deactivateDate = 0 : this.filterItems.CreatedAt = 0
      this.filter()
    },
    filterUsersBalance(balance) {
      let stat = 0
      this.filterItems[balance] !== 0 ? stat = -this.filterItems[balance] : stat = 1
      this.filterItems.BTCBalance = 0
      this.filterItems.EURBalance = 0
      this.filterItems.USDBalance = 0
      this.filterItems.USDTBalance = 0
      this.filterItems[balance] = stat
      this.filter()
    },
    checkingValues(cardElement) {
      if(cardElement === "" || typeof cardElement === undefined || cardElement === null) {
        return cardElement === "-"
      }
      if(cardElement === 0) {
        return cardElement === 0
      }
      return cardElement
    },
    filter() {
      this.filteredUsers = this.users.filter(user => {
        let cardInCycle = Object.values(user)
        for (const cardElement of cardInCycle) {
          this.checkingValues(cardElement)
        }
        const nickName = user.messenger.split(' ')[1]
        let country = this.filterItems.countries.find(item => item.code === user.country);
        let searchString = this.filterItems.search.toLowerCase()
        if ((!this.filterItems.role.name || user.role === this.filterItems.role.code)
            && (!this.filterItems.status.name || user.status === this.filterItems.status.code)
            && (user.name && user.email && nickName && country && user.lastIP)
            && (user.name.toLowerCase().indexOf(searchString) !== -1
                || user.email.toLowerCase().indexOf(searchString) !== -1
                || nickName.toLowerCase().indexOf(searchString) !== -1
                || country.name.toLowerCase().indexOf(searchString) !== -1
                || user.lastIP.indexOf(searchString)) !== -1) {
          return user
        }
      })
      if(this.filterItems.search && this.filteredUsers.length > 1) {
        this.background = "rgb(230, 255, 241)"
      } else {
        this.background = "none"
      }
      // фильтрация по дате
      if (this.filterItems.CreatedAt !== 0) {
        this.filterItems.CreatedAt === 1 ? this.filteredUsers.sort(function (a, b) {
          return new Date(b['CreatedAt']) - new Date(a['CreatedAt'])
        }) : this.filteredUsers.sort(function (a, b) {
          return new Date(a['CreatedAt']) - new Date(b['CreatedAt'])
        })
      }

      // фильтрация по балансам
      if (this.filterItems.USDTBalance !== 0) {
        this.filterItems.USDTBalance === 1 ? this.filteredUsers.sort(function (a, b) {
          return new Date(b['USDTBalance']) - new Date(a['USDTBalance'])
        }) : this.filteredUsers.sort(function (a, b) {
          return new Date(a['USDTBalance']) - new Date(b['USDTBalance'])
        })
      }

      if (this.filterItems.BTCBalance !== 0) {
        this.filterItems.BTCBalance === 1 ? this.filteredUsers.sort(function (a, b) {
          return new Date(b['BTCBalance']) - new Date(a['BTCBalance'])
        }) : this.filteredUsers.sort(function (a, b) {
          return new Date(a['BTCBalance']) - new Date(b['BTCBalance'])
        })
      }

      if (this.filterItems.USDBalance !== 0) {
        this.filterItems.USDBalance === 1 ? this.filteredUsers.sort(function (a, b) {
          return new Date(b['USDBalance']) - new Date(a['USDBalance'])
        }) : this.filteredUsers.sort(function (a, b) {
          return new Date(a['USDBalance']) - new Date(b['USDBalance'])
        })
      }

      if (this.filterItems.EURBalance !== 0) {
        this.filterItems.EURBalance === 1 ? this.filteredUsers.sort(function (a, b) {
          return new Date(b['EURBalance']) - new Date(a['EURBalance'])
        }) : this.filteredUsers.sort(function (a, b) {
          return new Date(a['EURBalance']) - new Date(b['EURBalance'])
        })
      }
    },
    resetFilterUsers() {
      this.filterItems = {
        CreatedAt: 1,
        countries: countries,
        statuses: [
          {
            name: "Active",
            code: 2
          },
          {
            name: "Moderation",
            code: 1
          },
          {
            name: "Unverified email",
            code: 0
          },
          {
            name: "Blocked",
            code: 3
          },
        ],
        status: {
          name: null,
          code: null
        },
        roles: [
          {
            name: "Administrator",
            code: 2
          },
          {
            name: "User",
            code: 1
          },
          {
            name: "Manager",
            code: 3
          },
          {
            name: "Invited",
            code: 0
          }
        ],
        role: {
          name: null,
          code: null
        },
        country: {
          name: null,
          code: null
        },
        search: '',
      }
      this.filter()
    },
    deleteUser(user) {
      this.$swal({
        title: "Вы действительно хотите удалить пользователя " + user.name + " " + user.email,
        type: "alert message",
        showCancelButton: true,
        confirmButtonClass: "btn-danger",
        confirmButtonText: "Да, удалить",
        icon: "error",
      }).then((result) => {
        if (result.value) {
          this.deleteUserMethod(user.ID)
              .catch(error => {
                console.log(error)
              });
        }
      })
    },
    deleteSelectUsers() {
      let selectUsers = []
      let IDs = []
      for (let [i, currentSelectUsers] in this.checkedUsers) {
        currentSelectUsers = this.checkedUsers[i]
        selectUsers.push(currentSelectUsers)
        for (const selectUser of selectUsers) {
          IDs.push(selectUser.ID)
        }
      }
      this.$swal({
        title: "Вы действительно хотите удалить выбранных пользователей",
        type: "alert message",
        showCancelButton: true,
        confirmButtonClass: "btn-danger",
        confirmButtonText: "Да, удалить",
        icon: "error",
      }).then((result) => {
        for (const ID of IDs) {
          if (result.value && IDs.length > 0) {
            this.deleteUserMethod(ID)
                .catch(error => {
                  console.log(error)
                });
          }
        }
      })
    },
    selectAllUsers() {
      if (this.check) {
        return this.checkedUsers = []
      }
      return [this.checkedUsers = this.filteredUsers, this.check]
    },
    showModal() {
      this.$refs.modal.openModal()
    },
    showEditModal(user) {
      this.$refs.edit.openModal(user)
    },
    showBalanceModal() {
      this.$refs.balance.openModal()
    },
    showBalanceModalUser(user, balance, nameBalance, currency) {
      this.$refs.balance.openModalUser(user, balance, nameBalance, currency)
    },
    getUsers() {
      if (this.loggedIn && this.users.length === 0 && this.req) {
        this.$store.dispatch("admin/getUsersMethod")
      }
      this.req = false
      return true
    },
    getCountry(userCountry) {
      let country = this.filterItems.countries.find(item => item.code === userCountry);
      if(country) {
        return country.name
      }
      return "-"
    },
    copySkype(myData) {
      if (!myData[1]) {
        let fadeInput = document.createElement("input");
        document.body.appendChild(fadeInput);
        fadeInput.setAttribute("id", "fadeInput");
        document.getElementById("fadeInput").value = myData[2];
        fadeInput.select();
        document.execCommand("copy");
        document.body.removeChild(fadeInput);
        this.$swal({
          title: "Nick " + myData[2] + " copied to clipboard",
          type: "alert message",
          icon: "info",
        })
        // timeout -> window.href
        setTimeout(() => window.open('https://web.skype.com/', '_blank'), 1000)
      } else {
        window.open('https://t.me/' + myData[2], '_blank')
      }
    },
    includeMessenger(user) {
      const tg = /TG/
      const skype = /SK/
      const userMessenger = user.messenger
      const nickName = userMessenger.split(' ')[1]
      if (tg.test(userMessenger) && user.messenger) {
        return ["https://t.me/" + nickName, "TG", nickName]
      }
      if (skype.test(userMessenger) && user.messenger) {
        return ["https://web.skype.com/", "SKYPE", nickName]
      }
      return ["", "", "",]
    },
    userBalance(balance) {
      if (balance === 0) {
        return '0.00'
      }

      let result = []
      let arrayBalance = ("" + balance).split("").map(Number)
      for (const i of arrayBalance) {
        result.push(i)
      }

      function roundingNum(x, n) {
        return parseFloat(Number.parseFloat(x).toFixed(n || 2));
      }

      let arrRoundingNum = Array.from(roundingNum(balance, 2).toString())
      let endResult = Array.from(arrRoundingNum.toString().split("."))[1]
      let firstResult = Array.from(arrRoundingNum.toString().split(".")[0])
      if (firstResult.toString().replace(/[\s.,%]/g, '').split('').length > 3) {
        arrRoundingNum.splice(0, 1)
        return result.shift() + ',' + arrRoundingNum.join("")
      }
      if (typeof endResult == 'undefined') {
        return arrRoundingNum.join("") + ".00"
      }
      if (endResult.split('').length < 2) {
        return arrRoundingNum.join("") + ".0"
      }

      return arrRoundingNum.join("")
    },
    splitDate(date) {
      if (date === "0001-01-01T00:00:00Z" || !date) {
        return "-"
      }
      let userTimeZone = this.currentUser.timeZone
      let timeZone = this.timezones.find(item => item.text === userTimeZone);
      if(timeZone && timeZone.utc.includes('-')) {
        let newDate = new Date(date);
        let utc = Number(timeZone.utc.split('-')[1])
        newDate.setTime(newDate.getTime() - (utc * 60 * 60 * 1000))
        return newDate.toLocaleString().split(",")[0] + " " + newDate.toLocaleString().split(",")[1]
      }
      if(timeZone && !timeZone.utc.includes('-')) {
        let newDate = new Date(date);
        let utc = Number(timeZone.utc)
        newDate.setTime(newDate.getTime() + (utc * 60 * 60 * 1000))
        return newDate.toLocaleString().split(",")[0] + " " + newDate.toLocaleString().split(",")[1]
      }
    },
    access(user) {
      this.$store.dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
          .then(() => {
            return apiV1.post("admin/getAccess", {email: user.email},
                {
                  headers: {
                    Authorization: this.getTokenHeader,
                    "Content-Type": "application/json"
                  },
                }
            )
                .then(response => {
                  if (!response.data.success) {
                    throw new Error(response.data.error);
                  }
                  window.location.href = process.env.VUE_APP_REDIRECT_URL + "/access/" + response.data.token;
                })
                .catch(error => {
                  console.log(error);
                  throw error;
                });
          });
    },
  },
  computed: {
    ...mapGetters({
      users: "admin/getUsers",
      loggedIn: "auth/isLoggedIn",
      getTokenHeader: "auth/getTokenHeader",
      currentUser: "auth/currentUser"
    })
  },
  components: {
    CreateUserModal,
    EditUserModal,
    EditBalanceModal
  }
}
</script>

<style scoped>
.table.gs-0 td:first-child {
  padding-left: 1rem !important;
}

.table th:last-child {
  padding: 0.75rem 0.75rem !important;
}

.show-delete {
  transition: opacity 0.2s ease !important;
}

:deep(.popper) {
  background: #eaeaea !important;
  padding: 7px 10px !important;
  border-radius: 5px !important;
  z-index: 999;
  color: var(--kt-text-dark);
  /*-webkit-box-shadow: 0px 2px 3px 1px rgba(0, 0, 0, 0.14) !important;*/
  /*-moz-box-shadow: 0px 2px 3px 1px rgba(0, 0, 0, 0.14) !important;*/
  /*box-shadow: 0px 4px 8px 4px rgba(0, 0, 0, 0.14) !important;*/
}

:deep(.popper #arrow::before) {
  background: #eaeaea !important;
}

:deep(.popper:hover),
:deep(.popper:hover > #arrow::before) {
  background: #eaeaea !important;
}
</style>