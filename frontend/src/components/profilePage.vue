<template>
  <div>
    <div class="d-flex flex-column flex-column-fluid container-fluid">
      <!--begin::Toolbar-->
      <div class="toolbar mb-5 mb-lg-7" id="kt_toolbar">
        <!--begin::Page title-->
        <div class="page-title d-flex flex-column me-3">
          <!--begin::Title-->
          <h1 class="d-flex text-dark fw-bold my-1 fs-3">{{ $t('profile.accountOverview') }}</h1>
          <!--end::Title-->
          <!--begin::Breadcrumb-->
          <ul class="breadcrumb breadcrumb-dot fw-semibold text-gray-600 fs-7 my-1">
            <!--begin::Item-->
            <li class="breadcrumb-item text-gray-600">
              <router-link to="/" class="text-gray-600 text-hover-primary">{{ $t('common.crumb1') }}</router-link>
            </li>
            <!--end::Item-->
            <!--begin::Item-->
            <li class="breadcrumb-item text-gray-600">{{ $t('profile.account') }}</li>
            <!--end::Item-->
            <!--begin::Item-->
            <li class="breadcrumb-item text-gray-500">{{ $t('common.settings') }}</li>
            <!--end::Item-->
          </ul>
          <!--end::Breadcrumb-->
        </div>
        <!--end::Page title-->
      </div>
      <!--end::Toolbar-->
      <!--begin::Post-->
      <div v-if="getUserInfo()" class="content flex-column-fluid">
        <!--begin::Navbar-->
        <div class="card mb-5 mb-xl-10">
          <div class="card-body pt-9 pb-0">
            <!--begin::Details-->
            <div class="d-flex flex-wrap flex-sm-nowrap mb-3">
              <!--begin: Pic-->
              <div class="me-7 mb-4">
                <div class="symbol symbol-100px symbol-lg-160px symbol-fixed position-relative">
                  <img :src="avatar"/>
                  <div
                      class="position-absolute translate-middle bottom-0 start-100 mb-6 bg-success rounded-circle border border-4 border-body h-20px w-20px"></div>
                </div>
              </div>
              <!--end::Pic-->
              <!--begin::Info-->
              <div class="flex-grow-1">
                <!--begin::Title-->
                <div class="d-flex justify-content-between align-items-start flex-wrap mb-2">
                  <!--begin::User-->
                  <div class="d-flex flex-column">
                    <!--begin::Name-->
                    <div class="d-flex align-items-center mb-2">
                      <a href="#"
                         class="text-gray-900 text-hover-primary fs-2 fw-bold me-1">{{ this.userinfo.name }}</a>
                      <a href="#">
                        <!--begin::Svg Icon | path: icons/duotune/general/gen026.svg-->
                        <span class="svg-icon svg-icon-1 svg-icon-primary">
																	<svg xmlns="http://www.w3.org/2000/svg" width="24px" height="24px"
                                       viewBox="0 0 24 24">
																		<path
                                        d="M10.0813 3.7242C10.8849 2.16438 13.1151 2.16438 13.9187 3.7242V3.7242C14.4016 4.66147 15.4909 5.1127 16.4951 4.79139V4.79139C18.1663 4.25668 19.7433 5.83365 19.2086 7.50485V7.50485C18.8873 8.50905 19.3385 9.59842 20.2758 10.0813V10.0813C21.8356 10.8849 21.8356 13.1151 20.2758 13.9187V13.9187C19.3385 14.4016 18.8873 15.491 19.2086 16.4951V16.4951C19.7433 18.1663 18.1663 19.7433 16.4951 19.2086V19.2086C15.491 18.8873 14.4016 19.3385 13.9187 20.2758V20.2758C13.1151 21.8356 10.8849 21.8356 10.0813 20.2758V20.2758C9.59842 19.3385 8.50905 18.8873 7.50485 19.2086V19.2086C5.83365 19.7433 4.25668 18.1663 4.79139 16.4951V16.4951C5.1127 15.491 4.66147 14.4016 3.7242 13.9187V13.9187C2.16438 13.1151 2.16438 10.8849 3.7242 10.0813V10.0813C4.66147 9.59842 5.1127 8.50905 4.79139 7.50485V7.50485C4.25668 5.83365 5.83365 4.25668 7.50485 4.79139V4.79139C8.50905 5.1127 9.59842 4.66147 10.0813 3.7242V3.7242Z"
                                        fill="currentColor"/>
																		<path
                                        d="M14.8563 9.1903C15.0606 8.94984 15.3771 8.9385 15.6175 9.14289C15.858 9.34728 15.8229 9.66433 15.6185 9.9048L11.863 14.6558C11.6554 14.9001 11.2876 14.9258 11.048 14.7128L8.47656 12.4271C8.24068 12.2174 8.21944 11.8563 8.42911 11.6204C8.63877 11.3845 8.99996 11.3633 9.23583 11.5729L11.3706 13.4705L14.8563 9.1903Z"
                                        fill="white"/>
																	</svg>
																</span>
                        <!--end::Svg Icon-->
                      </a>
                      <!--                      <a href="#" class="btn btn-sm btn-light-success fw-bold ms-2 fs-8 py-1 px-3" data-bs-toggle="modal" data-bs-target="#kt_modal_upgrade_plan">Upgrade to Pro</a>-->
                    </div>
                    <!--end::Name-->
                    <!--begin::Info-->
                    <div class="d-flex flex-wrap fw-semibold fs-6 mb-4 pe-2">
                      <a href="#" class="d-flex align-items-center text-gray-400 text-hover-primary me-5 mb-2">
                        <!--begin::Svg Icon | path: icons/duotune/communication/com006.svg-->
                        <span class="svg-icon svg-icon-4 me-1">
																<svg width="18" height="18" viewBox="0 0 18 18" fill="none"
                                     xmlns="http://www.w3.org/2000/svg">
																	<path opacity="0.3"
                                        d="M16.5 9C16.5 13.125 13.125 16.5 9 16.5C4.875 16.5 1.5 13.125 1.5 9C1.5 4.875 4.875 1.5 9 1.5C13.125 1.5 16.5 4.875 16.5 9Z"
                                        fill="currentColor"/>
																	<path
                                      d="M9 16.5C10.95 16.5 12.75 15.75 14.025 14.55C13.425 12.675 11.4 11.25 9 11.25C6.6 11.25 4.57499 12.675 3.97499 14.55C5.24999 15.75 7.05 16.5 9 16.5Z"
                                      fill="currentColor"/>
																	<rect x="7" y="6" width="4" height="4" rx="2" fill="currentColor"/>
																</svg>
															</span>
                        <!--end::Svg Icon-->
                        <span v-if="this.userinfo.role === 2">
                          Admin
                        </span>
                        <span v-if="this.userinfo.role === 1">
                          User
                        </span>
                        <span v-if="this.userinfo.role === 3">
                          Manager
                        </span>
                      </a>
                      <a href="#" class="d-flex align-items-center text-gray-400 text-hover-primary me-5 mb-2">
                        <!--begin::Svg Icon | path: icons/duotune/general/gen018.svg-->
                        <span class="svg-icon svg-icon-4 me-1">
																<svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                     xmlns="http://www.w3.org/2000/svg">
																	<path opacity="0.3"
                                        d="M18.0624 15.3453L13.1624 20.7453C12.5624 21.4453 11.5624 21.4453 10.9624 20.7453L6.06242 15.3453C4.56242 13.6453 3.76242 11.4453 4.06242 8.94534C4.56242 5.34534 7.46242 2.44534 11.0624 2.04534C15.8624 1.54534 19.9624 5.24534 19.9624 9.94534C20.0624 12.0453 19.2624 13.9453 18.0624 15.3453Z"
                                        fill="currentColor"/>
																	<path
                                      d="M12.0624 13.0453C13.7193 13.0453 15.0624 11.7022 15.0624 10.0453C15.0624 8.38849 13.7193 7.04535 12.0624 7.04535C10.4056 7.04535 9.06241 8.38849 9.06241 10.0453C9.06241 11.7022 10.4056 13.0453 12.0624 13.0453Z"
                                      fill="currentColor"/>
																</svg>
															</span>
                        <!--end::Svg Icon-->{{ getCountryToUser() }}</a>
                      <a href="#" class="d-flex align-items-center text-gray-400 text-hover-primary mb-2">
                        <!--begin::Svg Icon | path: icons/duotune/communication/com011.svg-->
                        <span class="svg-icon svg-icon-4 me-1">
																<svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                     xmlns="http://www.w3.org/2000/svg">
																	<path opacity="0.3"
                                        d="M21 19H3C2.4 19 2 18.6 2 18V6C2 5.4 2.4 5 3 5H21C21.6 5 22 5.4 22 6V18C22 18.6 21.6 19 21 19Z"
                                        fill="currentColor"/>
																	<path
                                      d="M21 5H2.99999C2.69999 5 2.49999 5.10005 2.29999 5.30005L11.2 13.3C11.7 13.7 12.4 13.7 12.8 13.3L21.7 5.30005C21.5 5.10005 21.3 5 21 5Z"
                                      fill="currentColor"/>
																</svg>
															</span>
                        <!--end::Svg Icon-->
                        {{ this.userinfo.email }}
                      </a>
                    </div>
                    <!--end::Info-->
                  </div>
                  <!--end::User-->
                  <!--begin::Actions-->
                  <!--                  <div class="d-flex my-4">-->
                  <!--                    <a href="#" class="btn btn-sm btn-light me-2" id="kt_user_follow_button">-->
                  <!--                      &lt;!&ndash;begin::Svg Icon | path: icons/duotune/arrows/arr012.svg&ndash;&gt;-->
                  <!--                      <span class="svg-icon svg-icon-3 d-none">-->
                  <!--																<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">-->
                  <!--																	<path opacity="0.3" d="M10 18C9.7 18 9.5 17.9 9.3 17.7L2.3 10.7C1.9 10.3 1.9 9.7 2.3 9.3C2.7 8.9 3.29999 8.9 3.69999 9.3L10.7 16.3C11.1 16.7 11.1 17.3 10.7 17.7C10.5 17.9 10.3 18 10 18Z" fill="currentColor" />-->
                  <!--																	<path d="M10 18C9.7 18 9.5 17.9 9.3 17.7C8.9 17.3 8.9 16.7 9.3 16.3L20.3 5.3C20.7 4.9 21.3 4.9 21.7 5.3C22.1 5.7 22.1 6.30002 21.7 6.70002L10.7 17.7C10.5 17.9 10.3 18 10 18Z" fill="currentColor" />-->
                  <!--																</svg>-->
                  <!--															</span>-->
                  <!--                      &lt;!&ndash;end::Svg Icon&ndash;&gt;-->
                  <!--                      &lt;!&ndash;begin::Indicator label&ndash;&gt;-->
                  <!--                      <span class="indicator-label">Follow</span>-->
                  <!--                      &lt;!&ndash;end::Indicator label&ndash;&gt;-->
                  <!--                      &lt;!&ndash;begin::Indicator progress&ndash;&gt;-->
                  <!--                      <span class="indicator-progress">Please wait...-->
                  <!--															<span class="spinner-border spinner-border-sm align-middle ms-2"></span></span>-->
                  <!--                      &lt;!&ndash;end::Indicator progress&ndash;&gt;-->
                  <!--                    </a>-->
                  <!--                    <a href="#" class="btn btn-sm btn-primary me-2" data-bs-toggle="modal" data-bs-target="#kt_modal_offer_a_deal">Hire Me</a>-->
                  <!--                    &lt;!&ndash;begin::Menu&ndash;&gt;-->
                  <!--                    <div class="me-0">-->
                  <!--                      <button class="btn btn-sm btn-icon btn-bg-light btn-active-color-primary" data-kt-menu-trigger="click" data-kt-menu-placement="bottom-end">-->
                  <!--                        <i class="bi bi-three-dots fs-3"></i>-->
                  <!--                      </button>-->
                  <!--                      &lt;!&ndash;begin::Menu 3&ndash;&gt;-->
                  <!--                      <div class="menu menu-sub menu-sub-dropdown menu-column menu-rounded menu-gray-800 menu-state-bg-light-primary fw-semibold w-200px py-3" data-kt-menu="true">-->
                  <!--                        &lt;!&ndash;begin::Heading&ndash;&gt;-->
                  <!--                        <div class="menu-item px-3">-->
                  <!--                          <div class="menu-content text-muted pb-2 px-3 fs-7 text-uppercase">Payments</div>-->
                  <!--                        </div>-->
                  <!--                        &lt;!&ndash;end::Heading&ndash;&gt;-->
                  <!--                        &lt;!&ndash;begin::Menu item&ndash;&gt;-->
                  <!--                        <div class="menu-item px-3">-->
                  <!--                          <a href="#" class="menu-link px-3">Create Invoice</a>-->
                  <!--                        </div>-->
                  <!--                        &lt;!&ndash;end::Menu item&ndash;&gt;-->
                  <!--                        &lt;!&ndash;begin::Menu item&ndash;&gt;-->
                  <!--                        <div class="menu-item px-3">-->
                  <!--                          <a href="#" class="menu-link flex-stack px-3">Create Payment-->
                  <!--                            <i class="fas fa-exclamation-circle ms-2 fs-7" data-bs-toggle="tooltip" title="Specify a target name for future usage and reference"></i></a>-->
                  <!--                        </div>-->
                  <!--                        &lt;!&ndash;end::Menu item&ndash;&gt;-->
                  <!--                        &lt;!&ndash;begin::Menu item&ndash;&gt;-->
                  <!--                        <div class="menu-item px-3">-->
                  <!--                          <a href="#" class="menu-link px-3">Generate Bill</a>-->
                  <!--                        </div>-->
                  <!--                        &lt;!&ndash;end::Menu item&ndash;&gt;-->
                  <!--                        &lt;!&ndash;begin::Menu item&ndash;&gt;-->
                  <!--                        <div class="menu-item px-3" data-kt-menu-trigger="hover" data-kt-menu-placement="right-end">-->
                  <!--                          <a href="#" class="menu-link px-3">-->
                  <!--                            <span class="menu-title">Subscription</span>-->
                  <!--                            <span class="menu-arrow"></span>-->
                  <!--                          </a>-->
                  <!--                          &lt;!&ndash;begin::Menu sub&ndash;&gt;-->
                  <!--                          <div class="menu-sub menu-sub-dropdown w-175px py-4">-->
                  <!--                            &lt;!&ndash;begin::Menu item&ndash;&gt;-->
                  <!--                            <div class="menu-item px-3">-->
                  <!--                              <a href="#" class="menu-link px-3">Plans</a>-->
                  <!--                            </div>-->
                  <!--                            &lt;!&ndash;end::Menu item&ndash;&gt;-->
                  <!--                            &lt;!&ndash;begin::Menu item&ndash;&gt;-->
                  <!--                            <div class="menu-item px-3">-->
                  <!--                              <a href="#" class="menu-link px-3">Billing</a>-->
                  <!--                            </div>-->
                  <!--                            &lt;!&ndash;end::Menu item&ndash;&gt;-->
                  <!--                            &lt;!&ndash;begin::Menu item&ndash;&gt;-->
                  <!--                            <div class="menu-item px-3">-->
                  <!--                              <a href="#" class="menu-link px-3">Statements</a>-->
                  <!--                            </div>-->
                  <!--                            &lt;!&ndash;end::Menu item&ndash;&gt;-->
                  <!--                            &lt;!&ndash;begin::Menu separator&ndash;&gt;-->
                  <!--                            <div class="separator my-2"></div>-->
                  <!--                            &lt;!&ndash;end::Menu separator&ndash;&gt;-->
                  <!--                            &lt;!&ndash;begin::Menu item&ndash;&gt;-->
                  <!--                            <div class="menu-item px-3">-->
                  <!--                              <div class="menu-content px-3">-->
                  <!--                                &lt;!&ndash;begin::Switch&ndash;&gt;-->
                  <!--                                <label class="form-check form-switch form-check-custom form-check-solid">-->
                  <!--                                  &lt;!&ndash;begin::Input&ndash;&gt;-->
                  <!--                                  <input class="form-check-input w-30px h-20px" type="checkbox" value="1" checked="checked" name="notifications" />-->
                  <!--                                  &lt;!&ndash;end::Input&ndash;&gt;-->
                  <!--                                  &lt;!&ndash;end::Label&ndash;&gt;-->
                  <!--                                  <span class="form-check-label text-muted fs-6">Recuring</span>-->
                  <!--                                  &lt;!&ndash;end::Label&ndash;&gt;-->
                  <!--                                </label>-->
                  <!--                                &lt;!&ndash;end::Switch&ndash;&gt;-->
                  <!--                              </div>-->
                  <!--                            </div>-->
                  <!--                            &lt;!&ndash;end::Menu item&ndash;&gt;-->
                  <!--                          </div>-->
                  <!--                          &lt;!&ndash;end::Menu sub&ndash;&gt;-->
                  <!--                        </div>-->
                  <!--                        &lt;!&ndash;end::Menu item&ndash;&gt;-->
                  <!--                        &lt;!&ndash;begin::Menu item&ndash;&gt;-->
                  <!--                        <div class="menu-item px-3 my-1">-->
                  <!--                          <a href="#" class="menu-link px-3">Settings</a>-->
                  <!--                        </div>-->
                  <!--                        &lt;!&ndash;end::Menu item&ndash;&gt;-->
                  <!--                      </div>-->
                  <!--                      &lt;!&ndash;end::Menu 3&ndash;&gt;-->
                  <!--                    </div>-->
                  <!--                    &lt;!&ndash;end::Menu&ndash;&gt;-->
                  <!--                  </div>-->
                  <!--end::Actions-->
                </div>
                <!--end::Title-->
                <!--begin::Stats-->
                <!--                <div class="d-flex flex-wrap flex-stack">-->
                <!--                  &lt;!&ndash;begin::Wrapper&ndash;&gt;-->
                <!--                  <div class="d-flex flex-column flex-grow-1 pe-8">-->
                <!--                    &lt;!&ndash;begin::Stats&ndash;&gt;-->
                <!--                    <div class="d-flex flex-wrap">-->
                <!--                      &lt;!&ndash;begin::Stat&ndash;&gt;-->
                <!--                      <div class="border border-gray-300 border-dashed rounded min-w-125px py-3 px-4 me-6 mb-3">-->
                <!--                        &lt;!&ndash;begin::Number&ndash;&gt;-->
                <!--                        <div class="d-flex align-items-center">-->
                <!--                          &lt;!&ndash;begin::Svg Icon | path: icons/duotune/arrows/arr066.svg&ndash;&gt;-->
                <!--                          <span class="svg-icon svg-icon-3 svg-icon-success me-2">-->
                <!--																		<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">-->
                <!--																			<rect opacity="0.5" x="13" y="6" width="13" height="2" rx="1" transform="rotate(90 13 6)" fill="currentColor" />-->
                <!--																			<path d="M12.5657 8.56569L16.75 12.75C17.1642 13.1642 17.8358 13.1642 18.25 12.75C18.6642 12.3358 18.6642 11.6642 18.25 11.25L12.7071 5.70711C12.3166 5.31658 11.6834 5.31658 11.2929 5.70711L5.75 11.25C5.33579 11.6642 5.33579 12.3358 5.75 12.75C6.16421 13.1642 6.83579 13.1642 7.25 12.75L11.4343 8.56569C11.7467 8.25327 12.2533 8.25327 12.5657 8.56569Z" fill="currentColor" />-->
                <!--																		</svg>-->
                <!--																	</span>-->
                <!--                          &lt;!&ndash;end::Svg Icon&ndash;&gt;-->
                <!--                          <div class="fs-2 fw-bold" data-kt-countup="true" data-kt-countup-value="4500" data-kt-countup-prefix="$">0</div>-->
                <!--                        </div>-->
                <!--                        &lt;!&ndash;end::Number&ndash;&gt;-->
                <!--                        &lt;!&ndash;begin::Label&ndash;&gt;-->
                <!--                        <div class="fw-semibold fs-6 text-gray-400">Earnings</div>-->
                <!--                        &lt;!&ndash;end::Label&ndash;&gt;-->
                <!--                      </div>-->
                <!--                      &lt;!&ndash;end::Stat&ndash;&gt;-->
                <!--                      &lt;!&ndash;begin::Stat&ndash;&gt;-->
                <!--                      <div class="border border-gray-300 border-dashed rounded min-w-125px py-3 px-4 me-6 mb-3">-->
                <!--                        &lt;!&ndash;begin::Number&ndash;&gt;-->
                <!--                        <div class="d-flex align-items-center">-->
                <!--                          &lt;!&ndash;begin::Svg Icon | path: icons/duotune/arrows/arr065.svg&ndash;&gt;-->
                <!--                          <span class="svg-icon svg-icon-3 svg-icon-danger me-2">-->
                <!--																		<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">-->
                <!--																			<rect opacity="0.5" x="11" y="18" width="13" height="2" rx="1" transform="rotate(-90 11 18)" fill="currentColor" />-->
                <!--																			<path d="M11.4343 15.4343L7.25 11.25C6.83579 10.8358 6.16421 10.8358 5.75 11.25C5.33579 11.6642 5.33579 12.3358 5.75 12.75L11.2929 18.2929C11.6834 18.6834 12.3166 18.6834 12.7071 18.2929L18.25 12.75C18.6642 12.3358 18.6642 11.6642 18.25 11.25C17.8358 10.8358 17.1642 10.8358 16.75 11.25L12.5657 15.4343C12.2533 15.7467 11.7467 15.7467 11.4343 15.4343Z" fill="currentColor" />-->
                <!--																		</svg>-->
                <!--																	</span>-->
                <!--                          &lt;!&ndash;end::Svg Icon&ndash;&gt;-->
                <!--                          <div class="fs-2 fw-bold" data-kt-countup="true" data-kt-countup-value="75">0</div>-->
                <!--                        </div>-->
                <!--                        &lt;!&ndash;end::Number&ndash;&gt;-->
                <!--                        &lt;!&ndash;begin::Label&ndash;&gt;-->
                <!--                        <div class="fw-semibold fs-6 text-gray-400">Projects</div>-->
                <!--                        &lt;!&ndash;end::Label&ndash;&gt;-->
                <!--                      </div>-->
                <!--                      &lt;!&ndash;end::Stat&ndash;&gt;-->
                <!--                      &lt;!&ndash;begin::Stat&ndash;&gt;-->
                <!--                      <div class="border border-gray-300 border-dashed rounded min-w-125px py-3 px-4 me-6 mb-3">-->
                <!--                        &lt;!&ndash;begin::Number&ndash;&gt;-->
                <!--                        <div class="d-flex align-items-center">-->
                <!--                          &lt;!&ndash;begin::Svg Icon | path: icons/duotune/arrows/arr066.svg&ndash;&gt;-->
                <!--                          <span class="svg-icon svg-icon-3 svg-icon-success me-2">-->
                <!--																		<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">-->
                <!--																			<rect opacity="0.5" x="13" y="6" width="13" height="2" rx="1" transform="rotate(90 13 6)" fill="currentColor" />-->
                <!--																			<path d="M12.5657 8.56569L16.75 12.75C17.1642 13.1642 17.8358 13.1642 18.25 12.75C18.6642 12.3358 18.6642 11.6642 18.25 11.25L12.7071 5.70711C12.3166 5.31658 11.6834 5.31658 11.2929 5.70711L5.75 11.25C5.33579 11.6642 5.33579 12.3358 5.75 12.75C6.16421 13.1642 6.83579 13.1642 7.25 12.75L11.4343 8.56569C11.7467 8.25327 12.2533 8.25327 12.5657 8.56569Z" fill="currentColor" />-->
                <!--																		</svg>-->
                <!--																	</span>-->
                <!--                          &lt;!&ndash;end::Svg Icon&ndash;&gt;-->
                <!--                          <div class="fs-2 fw-bold" data-kt-countup="true" data-kt-countup-value="60" data-kt-countup-prefix="%">0</div>-->
                <!--                        </div>-->
                <!--                        &lt;!&ndash;end::Number&ndash;&gt;-->
                <!--                        &lt;!&ndash;begin::Label&ndash;&gt;-->
                <!--                        <div class="fw-semibold fs-6 text-gray-400">Success Rate</div>-->
                <!--                        &lt;!&ndash;end::Label&ndash;&gt;-->
                <!--                      </div>-->
                <!--                      &lt;!&ndash;end::Stat&ndash;&gt;-->
                <!--                    </div>-->
                <!--                    &lt;!&ndash;end::Stats&ndash;&gt;-->
                <!--                  </div>-->
                <!--                  &lt;!&ndash;end::Wrapper&ndash;&gt;-->
                <!--                  &lt;!&ndash;begin::Progress&ndash;&gt;-->
                <!--                  <div class="d-flex align-items-center w-200px w-sm-300px flex-column mt-3">-->
                <!--                    <div class="d-flex justify-content-between w-100 mt-auto mb-2">-->
                <!--                      <span class="fw-semibold fs-6 text-gray-400">Profile Compleation</span>-->
                <!--                      <span class="fw-bold fs-6">50%</span>-->
                <!--                    </div>-->
                <!--                    <div class="h-5px mx-3 w-100 bg-light mb-3">-->
                <!--                      <div class="bg-success rounded h-5px" role="progressbar" style="width: 50%;" aria-valuenow="50" aria-valuemin="0" aria-valuemax="100"></div>-->
                <!--                    </div>-->
                <!--                  </div>-->
                <!--                  &lt;!&ndash;end::Progress&ndash;&gt;-->
                <!--                </div>-->
                <!--end::Stats-->
              </div>
              <!--end::Info-->
            </div>
            <!--end::Details-->
            <!--begin::Navs-->
            <ul class="nav nav-stretch nav-line-tabs nav-line-tabs-2x border-transparent fs-5 fw-bold">
              <!--begin::Nav item-->
              <li class="nav-item mt-2">
                <button class="nav-link text-active-primary ms-0 me-10 py-5 active">{{ $t('common.settings') }}</button>
              </li>
              <!--end::Nav item-->
              <!--begin::Nav item-->
              <!--              <li class="nav-item mt-2">-->
              <!--                <btn class="nav-link text-active-primary ms-0 me-10 py-5">{{$t('common.settings')}}</btn>-->
              <!--              </li>-->
              <!--              &lt;!&ndash;end::Nav item&ndash;&gt;-->
              <!--              &lt;!&ndash;begin::Nav item&ndash;&gt;-->
              <!--              <li class="nav-item mt-2">-->
              <!--                <a class="nav-link text-active-primary ms-0 me-10 py-5" href="../../demo14/dist/account/security.html">Security</a>-->
              <!--              </li>-->
              <!--              &lt;!&ndash;end::Nav item&ndash;&gt;-->
              <!--              &lt;!&ndash;begin::Nav item&ndash;&gt;-->
              <!--              <li class="nav-item mt-2">-->
              <!--                <a class="nav-link text-active-primary ms-0 me-10 py-5" href="../../demo14/dist/account/billing.html">Billing</a>-->
              <!--              </li>-->
              <!--              &lt;!&ndash;end::Nav item&ndash;&gt;-->
              <!--              &lt;!&ndash;begin::Nav item&ndash;&gt;-->
              <!--              <li class="nav-item mt-2">-->
              <!--                <a class="nav-link text-active-primary ms-0 me-10 py-5" href="../../demo14/dist/account/statements.html">Statements</a>-->
              <!--              </li>-->
              <!--              &lt;!&ndash;end::Nav item&ndash;&gt;-->
              <!--              &lt;!&ndash;begin::Nav item&ndash;&gt;-->
              <!--              <li class="nav-item mt-2">-->
              <!--                <a class="nav-link text-active-primary ms-0 me-10 py-5" href="../../demo14/dist/account/referrals.html">Referrals</a>-->
              <!--              </li>-->
              <!--              &lt;!&ndash;end::Nav item&ndash;&gt;-->
              <!--              &lt;!&ndash;begin::Nav item&ndash;&gt;-->
              <!--              <li class="nav-item mt-2">-->
              <!--                <a class="nav-link text-active-primary ms-0 me-10 py-5" href="../../demo14/dist/account/api-keys.html">API Keys</a>-->
              <!--              </li>-->
              <!--              &lt;!&ndash;end::Nav item&ndash;&gt;-->
              <!--              &lt;!&ndash;begin::Nav item&ndash;&gt;-->
              <!--              <li class="nav-item mt-2">-->
              <!--                <a class="nav-link text-active-primary ms-0 me-10 py-5" href="../../demo14/dist/account/logs.html">Logs</a>-->
              <!--              </li>-->
              <!--end::Nav item-->
            </ul>
            <!--begin::Navs-->
          </div>
        </div>
        <!--end::Navbar-->

        <div class="card mb-5 mb-xl-10">
          <!--begin::Card header-->
          <div class="card-header border-0 cursor-pointer" role="button" data-bs-toggle="collapse"
               data-bs-target="#kt_account_profile_details" aria-expanded="true"
               aria-controls="kt_account_profile_details">
            <!--begin::Card title-->
            <div class="card-title m-0">
              <h3 class="fw-bold m-0">{{ $t('profile.details') }}</h3>
            </div>
            <!--end::Card title-->
          </div>
          <!--begin::Card header-->
          <!--begin::Content-->
          <div id="kt_account_settings_profile_details" class="collapse show">
            <!--begin::Form-->
            <form id="kt_account_profile_details_form" class="form">
              <!--begin::Card body-->
              <div class="card-body border-top p-9">
                <!--begin::Input group-->
                <!--                <div class="row mb-6">-->
                <!--                  &lt;!&ndash;begin::Label&ndash;&gt;-->
                <!--                  <label class="col-lg-4 col-form-label fw-semibold fs-6">Avatar</label>-->
                <!--                  &lt;!&ndash;end::Label&ndash;&gt;-->
                <!--                  &lt;!&ndash;begin::Col&ndash;&gt;-->
                <!--                  <div class="col-lg-8">-->
                <!--                    &lt;!&ndash;begin::Image input&ndash;&gt;-->
                <!--                    <div class="image-input image-input-outline" data-kt-image-input="true"-->
                <!--                         style="background-image: url('assets/media/svg/avatars/blank.svg')">-->
                <!--                      &lt;!&ndash;begin::Preview existing avatar&ndash;&gt;-->
                <!--                      <img :src="avatar" class="image-input-wrapper w-125px h-125px">-->

                <!--                      &lt;!&ndash;end::Preview existing avatar&ndash;&gt;-->
                <!--                      &lt;!&ndash;begin::Label&ndash;&gt;-->
                <!--                      <label class="btn btn-icon btn-circle btn-active-color-primary w-25px h-25px bg-body shadow"-->
                <!--                             data-kt-image-input-action="change" data-bs-toggle="tooltip" title="Change avatar">-->
                <!--                        <i class="bi bi-pencil-fill fs-7"></i>-->
                <!--                        &lt;!&ndash;begin::Inputs&ndash;&gt;-->
                <!--                        <input type="file" name="avatar" accept=".png, .jpg, .jpeg"/>-->
                <!--                        <input type="hidden" name="avatar_remove"/>-->
                <!--                        &lt;!&ndash;end::Inputs&ndash;&gt;-->
                <!--                      </label>-->
                <!--                      &lt;!&ndash;end::Label&ndash;&gt;-->
                <!--                      &lt;!&ndash;begin::Cancel&ndash;&gt;-->
                <!--                      <span class="btn btn-icon btn-circle btn-active-color-primary w-25px h-25px bg-body shadow"-->
                <!--                            data-kt-image-input-action="cancel" data-bs-toggle="tooltip" title="Cancel avatar">-->
                <!--																<i class="bi bi-x fs-2"></i>-->
                <!--															</span>-->
                <!--                      &lt;!&ndash;end::Cancel&ndash;&gt;-->
                <!--                      &lt;!&ndash;begin::Remove&ndash;&gt;-->
                <!--                      <span class="btn btn-icon btn-circle btn-active-color-primary w-25px h-25px bg-body shadow"-->
                <!--                            data-kt-image-input-action="remove" data-bs-toggle="tooltip" title="Remove avatar">-->
                <!--																<i class="bi bi-x fs-2"></i>-->
                <!--															</span>-->
                <!--                      &lt;!&ndash;end::Remove&ndash;&gt;-->
                <!--                    </div>-->
                <!--                    &lt;!&ndash;end::Image input&ndash;&gt;-->
                <!--                    &lt;!&ndash;begin::Hint&ndash;&gt;-->
                <!--                    <div class="form-text">Allowed file types: png, jpg, jpeg.</div>-->
                <!--                    &lt;!&ndash;end::Hint&ndash;&gt;-->
                <!--                  </div>-->
                <!--                  &lt;!&ndash;end::Col&ndash;&gt;-->
                <!--                </div>-->
                <!--end::Input group-->
                <!--begin::Input group-->
                <div class="row mb-6">
                  <!--begin::Label-->
                  <label class="col-lg-4 col-form-label required fw-semibold fs-6">{{ $t('common.fullName') }}</label>
                  <!--end::Label-->
                  <!--begin::Col-->
                  <div class="col-lg-8">
                    <!--begin::Row-->
                    <div class="row">
                      <!--begin::Col-->
                      <div class="col-lg-6 fv-row">
                        <input v-model="form.name" type="text" name="fname"
                               class="form-control form-control-lg form-control-solid mb-3 mb-lg-0"
                               :placeholder="$t('common.fullName') + '..'"/>
                      </div>
                      <!--end::Col-->
                      <!--begin::Col-->
                      <div class="col-lg-6 fv-row">
                        <input type="text" name="lname" class="form-control form-control-lg form-control-solid"
                               placeholder="Last name" value=""/>
                      </div>
                      <!--end::Col-->
                    </div>
                    <!--end::Row-->
                  </div>
                  <!--end::Col-->
                </div>
                <!--end::Input group-->
                <!--begin::Input group-->
                <div class="row mb-6">
                  <!--begin::Label-->
                  <label class="col-lg-4 col-form-label fw-semibold fs-6">
                    <span class="required">{{ $t('signUp.country') }}</span>
                    <i class="fas fa-exclamation-circle ms-1 fs-7" data-bs-toggle="tooltip"
                       title="Country of origination"></i>
                  </label>
                  <!--end::Label-->
                  <!--begin::Col-->
                  <div class="col-lg-8 fv-row select2-group">
                    <Select2 id="country" name="country" modelValue="countryForSend" v-model="countryForSend"
                             :options="countriesProfile"
                             @change="myChangeEvent($event)" @select="mySelectEvent($event)"/>
                    <select name="country" :placeholder="$t('signUp.country') + '..'"
                            class="form-select form-select-solid form-select-lg fw-semibold">
                      <option :value=this.countryForSend>{{ this.countryForSend.text }}</option>
                    </select>
                  </div>
                  <!--end::Col-->
                </div>
                <!--end::Input group-->
                <!--begin::Input group-->
                <div class="row mb-6">
                  <!--begin::Label-->
                  <label class="col-lg-4 col-form-label required fw-semibold fs-6">{{ $t('profile.timezone') }}</label>
                  <!--end::Label-->
                  <!--begin::Col-->
                  <div class="col-lg-8 fv-row select2-group">
                    <Select2 id="timezone" name="timezone" modelValue="timeZone" v-model="timeZone" :options="timeZones"
                             @change="myChangeEventTimeZone($event)" @select="mySelectEventTimeZone($event)"/>
                    <select name="timezone" :placeholder="$t('common.timeZone') + '..'"
                            class="form-select form-select-solid form-select-lg">
                      <option :value="this.timeZone">{{ this.timeZone.text }}</option>
                    </select>
                  </div>
                  <!--end::Col-->
                </div>
                <!--end::Input group-->
                <!--begin::Input group-->
                <div class="row mb-6">
                  <!--begin::Label-->
                  <label class="col-lg-4 col-form-label fw-semibold fs-6">
                    <span class="required">{{ $t('profile.contactDetails') }}</span>
                    <i class="fas fa-exclamation-circle ms-1 fs-7" data-bs-toggle="tooltip"
                       title="Phone number must be active"></i>
                  </label>
                  <!--end::Label-->
                  <!--begin::Col-->
                  <div class="col-lg-8 fv-row mb-6">
                    <div class="row">
                      <div class="col-lg-6 fv-row mb-3 mb-lg-0">
                        <Select2 id="messenger" name="messenger" modelValue="messenger" v-model="messenger"
                                 :options="messengers"
                                 @change="myChangeEventMessenger($event)" @select="mySelectEventMessenger($event)"/>
                        <select name="messenger" :placeholder="$t('signUp.messenger') + '..'"
                                class="form-select form-select-solid form-select-lg">
                          <option :value="this.messenger">{{ this.messenger.text }}</option>
                        </select>
                      </div>
                      <div class="col-lg-6 fv-row">
                        <input v-model="form.messengerNickname" type="text" name="nickName"
                               class="form-control form-control-lg form-control-solid"
                               :placeholder="$t('signUp.messengerNickname') + '..'"/>
                      </div>
                    </div>
                  </div>
                  <!--end::Col-->
                </div>
                <!--end::Input group-->
                <div class="d-flex align-items-center">
                  <span class="fs-6 fw-semibold">
                    {{ $t('profile.changeEmail') }}
                  </span>
                  <div class="d-flex ml-2 align-items-center">
                    <a @click="copySkype" class="copy-btn mr-2 d-flex align-items-center cursor-pointer">
                      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="35px" height="35px">
                        <path fill="#03a9f4" d="M24 4A20 20 0 1 0 24 44A20 20 0 1 0 24 4Z"/>
                        <path fill="#03a9f4"
                              d="M33.5 22A11.5 11.5 0 1 0 33.5 45 11.5 11.5 0 1 0 33.5 22zM14.5 3A11.5 11.5 0 1 0 14.5 26 11.5 11.5 0 1 0 14.5 3z"/>
                        <path fill="#fff"
                              d="M24.602,36C18,36,15,32.699,15,30.199C15,28.898,15.898,28,17.199,28c2.801,0,2.102,4.102,7.402,4.102c2.699,0,4.199-1.5,4.199-3c0-0.902-0.402-1.902-2.199-2.402l-5.902-1.5C16,24,15.102,21.398,15.102,18.898c0-5.098,4.699-6.898,9.098-6.898C28.301,12,33,14.199,33,17.199c0,1.301-1,2.102-2.301,2.102c-2.398,0-2-3.402-6.801-3.402c-2.398,0-3.797,1.102-3.797,2.703c0,1.598,1.898,2.098,3.598,2.5l4.402,1C32.898,23.199,34,26,34,28.699C33.898,32.898,30.898,36,24.602,36z"/>
                      </svg>
                    </a>
                    <a href="https://t.me/roman_gratscards" target="_blank"
                       class="copy-btn d-flex align-items-center cursor-pointer">
                      <svg width="35" height="35" viewBox="0 0 23 23" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path
                            d="M11.5 23C17.8513 23 23 17.8513 23 11.5C23 5.14873 17.8513 0 11.5 0C5.14873 0 0 5.14873 0 11.5C0 17.8513 5.14873 23 11.5 23Z"
                            fill="#039BE5"/>
                        <path
                            d="M5.26234 11.2509L16.3503 6.97578C16.8649 6.78987 17.3143 7.10133 17.1476 7.87949L17.1486 7.87854L15.2606 16.7728C15.1207 17.4034 14.746 17.5567 14.2218 17.2597L11.3468 15.1408L9.96009 16.4767C9.80676 16.63 9.67738 16.7594 9.3803 16.7594L9.58442 13.8336L14.9128 9.01991C15.1447 8.81579 14.861 8.70079 14.5553 8.90395L7.97059 13.0497L5.132 12.1642C4.5158 11.9687 4.50238 11.548 5.26234 11.2509Z"
                            fill="white"/>
                      </svg>
                    </a>
                  </div>
                </div>
              </div>
              <!--end::Card body-->
              <!--begin::Actions-->
              <div class="card-footer d-flex justify-content-end py-6 px-9">
                <button type="reset" class="btn btn-light btn-active-light-primary me-2">{{ $t('common.discard') }}
                </button>
                <button @click="checkForm" type="submit" class="btn btn-primary">{{ $t('profile.saveChanges') }}
                </button>
              </div>
              <!--end::Actions-->
            </form>
            <!--end::Form-->
          </div>
          <!--end::Content-->
        </div>
        <div class="card mb-5 mb-xl-10">
          <!--begin::Card header-->
          <div class="card-header border-0 cursor-pointer" role="button" data-bs-toggle="collapse"
               data-bs-target="#kt_account_signin_method">
            <div class="card-title m-0">
              <h3 class="fw-bold m-0">Sign-in Method</h3>
            </div>
          </div>
          <!--end::Card header-->
          <!--begin::Content-->
          <div id="kt_account_settings_signin_method" class="collapse show">
            <!--begin::Card body-->
            <div class="card-body border-top p-9">
              <!--              &lt;!&ndash;begin::Email Address&ndash;&gt;-->
              <!--              <div class="d-flex flex-wrap align-items-center">-->
              <!--                &lt;!&ndash;begin::Label&ndash;&gt;-->
              <!--                <div id="kt_signin_email">-->
              <!--                  <div class="fs-6 fw-bold mb-1">Email Address</div>-->
              <!--                  <div class="fw-semibold text-gray-600">support@keenthemes.com</div>-->
              <!--                </div>-->
              <!--                &lt;!&ndash;end::Label&ndash;&gt;-->
              <!--                &lt;!&ndash;begin::Edit&ndash;&gt;-->
              <!--                <div id="kt_signin_email_edit" class="flex-row-fluid d-none">-->
              <!--                  &lt;!&ndash;begin::Form&ndash;&gt;-->
              <!--                  <form id="kt_signin_change_email" class="form" novalidate="novalidate">-->
              <!--                    <div class="row mb-6">-->
              <!--                      <div class="col-lg-6 mb-4 mb-lg-0">-->
              <!--                        <div class="fv-row mb-0">-->
              <!--                          <label for="emailaddress" class="form-label fs-6 fw-bold mb-3">Enter New Email Address</label>-->
              <!--                          <input type="email" class="form-control form-control-lg form-control-solid" id="emailaddress"-->
              <!--                                 placeholder="Email Address" name="emailaddress" value="support@keenthemes.com"/>-->
              <!--                        </div>-->
              <!--                      </div>-->
              <!--                      <div class="col-lg-6">-->
              <!--                        <div class="fv-row mb-0">-->
              <!--                          <label for="confirmemailpassword" class="form-label fs-6 fw-bold mb-3">Confirm-->
              <!--                            Password</label>-->
              <!--                          <input type="password" class="form-control form-control-lg form-control-solid"-->
              <!--                                 name="confirmemailpassword" id="confirmemailpassword"/>-->
              <!--                        </div>-->
              <!--                      </div>-->
              <!--                    </div>-->
              <!--                    <div class="d-flex">-->
              <!--                      <button id="kt_signin_submit" type="button" class="btn btn-primary me-2 px-6">Update Email-->
              <!--                      </button>-->
              <!--                      <button id="kt_signin_cancel" type="button"-->
              <!--                              class="btn btn-color-gray-400 btn-active-light-primary px-6">Cancel-->
              <!--                      </button>-->
              <!--                    </div>-->
              <!--                  </form>-->
              <!--                  &lt;!&ndash;end::Form&ndash;&gt;-->
              <!--                </div>-->
              <!--                &lt;!&ndash;end::Edit&ndash;&gt;-->
              <!--                &lt;!&ndash;begin::Action&ndash;&gt;-->
              <!--                <div id="kt_signin_email_button" class="ms-auto">-->
              <!--                  <button disabled class="btn btn-light btn-active-light-primary">Change Email</button>-->
              <!--                </div>-->
              <!--                &lt;!&ndash;end::Action&ndash;&gt;-->
              <!--              </div>-->
              <!--              &lt;!&ndash;end::Email Address&ndash;&gt;-->
              <!--              &lt;!&ndash;begin::Separator&ndash;&gt;-->
              <!--              <div class="separator separator-dashed my-6"></div>-->
              <!--              &lt;!&ndash;end::Separator&ndash;&gt;-->
              <!--begin::Password-->
              <div class="d-flex flex-wrap align-items-center">
                <!--begin::Label-->
                <div v-if="!this.resetPassword">
                  <div class="fs-6 fw-bold mb-1">{{ $t('common.password') }}</div>
                  <div class="fw-semibold text-gray-600">************</div>
                </div>
                <form class="w-100 d-flex justify-content-between position-relative" v-else>
                  <div class="col-6 form-group fv-row mb-4">
                    <label for="oldPassword"
                           class="required fs-6 fw-semibold mb-2">{{ $t('profile.oldPassword') }}</label>
                    <input name="oldPassword" type="password" :placeholder="$t('profile.oldPassword') + '...'"
                           v-model="formPassword.oldPassword" class="form-control form-control-lg form-control-solid"
                           required/>
                  </div>
                  <div class="col-6 form-group fv-row">
                    <label for="newPassword" class="required fs-6 fw-semibold mb-2">{{
                        $t('profile.newPassword')
                      }}</label>
                    <input name="newPassword" type="password" :placeholder="$t('profile.newPassword') + '...'"
                           v-model="formPassword.password" class="form-control form-control-lg form-control-solid"
                           required/>
                  </div>
                  <p class="hidden-error-block password-error" v-if="!errors.password"
                     style="position: absolute;padding: 0px;top: 7rem;left: 0;right: 0;text-align: center;color: red;">
                    {{ this.errorText }}</p>
                </form>
                <!--end::Label-->
                <!--begin::Edit-->
                <div id="kt_signin_password_edit" class="flex-row-fluid d-none">
                  <!--begin::Form-->
                  <form id="kt_signin_change_password" class="form" novalidate="novalidate">
                    <div class="row mb-1">
                      <div class="col-lg-4">
                        <div class="fv-row mb-0">
                          <label for="currentpassword" class="form-label fs-6 fw-bold mb-3">Current Password</label>
                          <input type="password" class="form-control form-control-lg form-control-solid"
                                 name="currentpassword" id="currentpassword"/>
                        </div>
                      </div>
                      <div class="col-lg-4">
                        <div class="fv-row mb-0">
                          <label for="newpassword" class="form-label fs-6 fw-bold mb-3">New Password</label>
                          <input type="password" class="form-control form-control-lg form-control-solid"
                                 name="newpassword" id="newpassword"/>
                        </div>
                      </div>
                      <div class="col-lg-4">
                        <div class="fv-row mb-0">
                          <label for="confirmpassword" class="form-label fs-6 fw-bold mb-3">Confirm New Password</label>
                          <input type="password" class="form-control form-control-lg form-control-solid"
                                 name="confirmpassword" id="confirmpassword"/>
                        </div>
                      </div>
                    </div>
                    <div class="form-text mb-5">Password must be at least 8 character and contain symbols</div>
                    <div class="d-flex">
                      <button id="kt_password_submit" type="button" class="btn btn-primary me-2 px-6">Update Password
                      </button>
                      <button id="kt_password_cancel" type="button"
                              class="btn btn-color-gray-400 btn-active-light-primary px-6">Cancel
                      </button>
                    </div>
                  </form>
                  <!--end::Form-->
                </div>
                <!--end::Edit-->
                <!--begin::Action-->
                <div v-if="!this.resetPassword" class="ms-auto">
                  <button @click="this.resetPassword = true" class="btn btn-light btn-active-light-primary">
                    {{ $t('profile.resetPassword') }}
                  </button>
                </div>
                <!--end::Action-->
              </div>
              <!--end::Password-->
              <!--              &lt;!&ndash;begin::Notice&ndash;&gt;-->
              <!--              <div class="notice d-flex bg-light-primary rounded border-primary border border-dashed p-6">-->
              <!--                &lt;!&ndash;begin::Icon&ndash;&gt;-->
              <!--                &lt;!&ndash;begin::Svg Icon | path: icons/duotune/general/gen048.svg&ndash;&gt;-->
              <!--                <span class="svg-icon svg-icon-2tx svg-icon-primary me-4">-->
              <!--													<svg width="24" height="24" viewBox="0 0 24 24" fill="none"-->
              <!--                               xmlns="http://www.w3.org/2000/svg">-->
              <!--														<path opacity="0.3"-->
              <!--                                  d="M20.5543 4.37824L12.1798 2.02473C12.0626 1.99176 11.9376 1.99176 11.8203 2.02473L3.44572 4.37824C3.18118 4.45258 3 4.6807 3 4.93945V13.569C3 14.6914 3.48509 15.8404 4.4417 16.984C5.17231 17.8575 6.18314 18.7345 7.446 19.5909C9.56752 21.0295 11.6566 21.912 11.7445 21.9488C11.8258 21.9829 11.9129 22 12.0001 22C12.0872 22 12.1744 21.983 12.2557 21.9488C12.3435 21.912 14.4326 21.0295 16.5541 19.5909C17.8169 18.7345 18.8277 17.8575 19.5584 16.984C20.515 15.8404 21 14.6914 21 13.569V4.93945C21 4.6807 20.8189 4.45258 20.5543 4.37824Z"-->
              <!--                                  fill="currentColor"/>-->
              <!--														<path-->
              <!--                                d="M10.5606 11.3042L9.57283 10.3018C9.28174 10.0065 8.80522 10.0065 8.51412 10.3018C8.22897 10.5912 8.22897 11.0559 8.51412 11.3452L10.4182 13.2773C10.8099 13.6747 11.451 13.6747 11.8427 13.2773L15.4859 9.58051C15.771 9.29117 15.771 8.82648 15.4859 8.53714C15.1948 8.24176 14.7183 8.24176 14.4272 8.53714L11.7002 11.3042C11.3869 11.6221 10.874 11.6221 10.5606 11.3042Z"-->
              <!--                                fill="currentColor"/>-->
              <!--													</svg>-->
              <!--												</span>-->
              <!--                &lt;!&ndash;end::Svg Icon&ndash;&gt;-->
              <!--                &lt;!&ndash;end::Icon&ndash;&gt;-->
              <!--                &lt;!&ndash;begin::Wrapper&ndash;&gt;-->
              <!--                <div class="d-flex flex-stack flex-grow-1 flex-wrap flex-md-nowrap">-->
              <!--                  &lt;!&ndash;begin::Content&ndash;&gt;-->
              <!--                  <div class="mb-3 mb-md-0 fw-semibold">-->
              <!--                    <h4 class="text-gray-900 fw-bold">Secure Your Account</h4>-->
              <!--                    <div class="fs-6 text-gray-700 pe-7">Two-factor authentication adds an extra layer of security to-->
              <!--                      your account. To log in, in addition you'll need to provide a 6 digit code-->
              <!--                    </div>-->
              <!--                  </div>-->
              <!--                  &lt;!&ndash;end::Content&ndash;&gt;-->
              <!--                  &lt;!&ndash;begin::Action&ndash;&gt;-->
              <!--                  <a href="#" class="btn btn-primary px-6 align-self-center text-nowrap" data-bs-toggle="modal"-->
              <!--                     data-bs-target="#kt_modal_two_factor_authentication">Enable</a>-->
              <!--                  &lt;!&ndash;end::Action&ndash;&gt;-->
              <!--                </div>-->
              <!--                &lt;!&ndash;end::Wrapper&ndash;&gt;-->
              <!--              </div>-->
              <!--              &lt;!&ndash;end::Notice&ndash;&gt;-->
            </div>
            <!--end::Card body-->
            <div v-if="this.resetPassword" class="card-footer d-flex justify-content-center py-6 px-9">
              <button @click="cancelResetPassword" class="btn btn-light btn-active-light-primary">
                {{ $t('common.cancel') }}
              </button>
              <button :disabled="this.disabled" @click="sendResetPassword" class="btn btn-primary btn-active-primary">
                {{ $t('common.save') }}
              </button>
            </div>
          </div>
          <!--end::Content-->
        </div>
      </div>
    </div>
  </div>
  <!--  <button type="button" class="btn btn-danger" v-if="!this.loggedIn" @click="$router.push('/login')">{{ $t('commonErrors.notSignedIn') }}</button>-->
</template>

<script>
import {defineComponent} from 'vue'
import {mapActions, mapGetters} from "vuex";
import countriesProfile from '../countries-profile.json'
import timeZones from '../timezones.json'
import countries from '../countries.json'
import avatar from '@/assets/media/001-boy.svg'
import Select2 from 'vue3-select2-component';

export default defineComponent({
  name: 'profilePage',
  data() {
    return {
      resetPassword: false,
      avatar,
      req: true,
      countryForSend: "",
      timeZone: "",
      messenger: "",
      country: null,
      disabled: true,
      countries: countries,
      countriesProfile: countriesProfile,
      timeZones: timeZones,
      errorText: null,
      messengers: [{"id": "TG", "text": "Telegram"}, {"id": "SK", "text": "Skype"}],
      errors: {
        password: false,
        responseError: null,
      },
      form: {
        country: null,
        timeZone: null,
        messenger: null,
        messengerNickname: null,
        name: null
      },
      formPassword: {
        type: "Password",
        oldPassword: null,
        password: null
      }
    }
  },
  components: {Select2},
  watch: {
    userinfo: "getData",
    'formPassword.password': 'checkPasswords',
    'formPassword.oldPassword': 'checkPasswords',
  },
  created() {
    if (this.userinfo) {
      this.getData()
    }
  },
  methods: {
    ...mapActions({
      changeInfo: "users/changeUserInfo"
    }),
    checkPasswords: function () {
      this.errors.password = this.formPassword.password !== this.formPassword.oldPassword
      if (this.errors.password) {
        this.errorText = this.$t('profile.checkPassword')
        this.disabled = true
      }
      if (this.formPassword.password && this.formPassword.password.length < 6) {
        this.errorText = this.$t('profile.lengthPassword')
        this.disabled = true
      }
      if (this.formPassword.oldPassword && this.formPassword.oldPassword.length < 6) {
        this.errorText = this.$t('profile.lengthPassword')
        this.disabled = true
      }
      if (!this.formPassword.oldPassword && !this.formPassword.password) {
        this.disabled = true
      }
      if (this.errors.password && this.formPassword.oldPassword && (this.formPassword.oldPassword.length > 6) && this.formPassword.password && (this.formPassword.password.length > 6)) {
        this.disabled = false
      }
    },
    getCountryToUser() {
      let country = this.countries.find(item => item.code === this.userinfo.country);
      if (country) {
        this.country = country
        return country.code + ", " + country.name
      }
      return "-"
    },
    getData() {
      this.form.name = this.userinfo.name
      let country = this.countriesProfile.find(item => item.id === this.userinfo.country);
      if (country) {
        this.countryForSend = {
          id: country.id,
          text: country.text
        }
      }
      let timezone = this.timeZones.find(item => item.text === this.userinfo.timeZone);
      if (timezone) {
        this.timeZone = {
          id: timezone.id,
          text: timezone.text,
          utc: timezone.utc
        }
      }
      const tg = /TG/
      const skype = /SK/
      if (this.userinfo.messenger) {
        this.form.messengerNickname = this.userinfo.messenger.split(" ")[1]
      }
      if (this.userinfo.messenger && tg.test(this.userinfo.messenger)) {
        this.messenger = {
          id: "TG",
          text: "Telegram"
        }
      }
      if (this.userinfo.messenger && skype.test(this.userinfo.messenger)) {
        this.messenger = {
          id: "SK",
          text: "Skype"
        }
      }
    },
    getUserInfo() {
      if (this.loggedIn && !this.userInfoLoaded && this.req) {
        this.$store.dispatch("users/getUserInfoMethod");
      }
      this.req = false
      return true
    },
    myChangeEvent(val) {
      return val
    },
    myChangeEventTimeZone(val) {
      return val
    },
    mySelectEvent({id, text}) {
      this.countryForSend = {id, text}
      return this.countryForSend
    },
    mySelectEventTimeZone({id, text}) {
      this.timeZone = {id, text}
      return this.timeZone
    },
    myChangeEventMessenger(val) {
      return val
    },
    mySelectEventMessenger({id, text}) {
      this.messenger = {id, text}
      return this.messenger
    },
    copySkype() {
      let fadeInput = document.createElement("input");
      document.body.appendChild(fadeInput);
      fadeInput.setAttribute("id", "fadeInput");
      document.getElementById("fadeInput").value = "live:.cid.a4da24ca1a9a7857";
      fadeInput.select();
      document.execCommand("copy");
      document.body.removeChild(fadeInput);
      this.$swal({
        title: this.$t('common.copyNick'),
        type: "alert message",
        icon: "info",
      })
      // timeout -> window.href
      setTimeout(() => window.open('https://web.skype.com/', '_blank'), 1000)
    },
    checkForm: function (e) {
      e.preventDefault()
      this.form.country = this.countryForSend.id
      this.form.timeZone = this.timeZone.text
      this.form.messenger = this.messenger.id
    },
    cancelResetPassword() {
      this.resetPassword = false
      this.formPassword.password = null
    },
    sendResetPassword() {
      console.log(this.formPassword)
      this.changeInfo(this.formPassword)
          .then(() => {
            this.resetPassword = false
            this.$swal({
              title: this.$t('profile.changePasswordSuccess'),
              confirmButtonColor: "#50cd89",
              type: "alert message",
              icon: "success",
            })
          })
          .catch(error => {
            console.log(error)
            this.errors.responseError = error
            if (this.errors.responseError.toString() === "Error: badRequest") this.errorTitle = "Something went wrong. Try again later."
            if (this.errors.responseError.toString() === "Error: userNotFound") this.errorTitle = "User with this email address is not found."
            if (this.errors.responseError.toString() === "Error: registrationNotCompleted") this.errorTitle = "This user is not completely registered. Check your email for any further instructions."
            if (this.errors.responseError.toString() === "Error: incorrectPassword") this.errorTitle = "Incorrect password."
            console.log(this.errorTitle)
            this.$swal({
              title: this.errorTitle,
              type: "alert message",
              icon: "error",
              confirmButtonColor: "#f1416c",
            })
            this.disabled = false
          })
    }
  },
  computed: {
    ...mapGetters({
      userinfo: "users/getUserInfo",
      userInfoLoaded: "users/getUserInfoLoaded",
      loggedIn: "auth/isLoggedIn",
    }),
  }
})
</script>

<style>
.select2-group {
  position: relative;
}

.select2-container--focus .select2-selection, .select2-container--open .select2-selection {
  background-color: var(--kt-input-solid-bg-focus);
}

.select2-container.select2-selection {
  background-color: var(--kt-input-solid-bg);
  border-color: var(--kt-input-solid-bg);
  color: var(--kt-input-solid-color);
  transition: color 0.2s ease;
}

.select2-container .select2-selection--single {
  display: flex;
  align-items: center;
}

.select2-container .select2-selection {
  height: 50px;
  opacity: 0;
  box-shadow: none !important;
  outline: none !important;
}

.select2-container {
  box-sizing: border-box;
  display: inline-block;
  margin: 0;
  position: absolute;
  vertical-align: middle;
}

.select2-container .select2-dropdown {
  border: 0;
  box-shadow: var(--kt-dropdown-box-shadow);
  border-radius: 0.475rem;
  padding: 1rem 0;
  background-color: var(--kt-dropdown-bg);
}

.select2-dropdown .select2-search .select2-search__field {
  background-color: var(--kt-body-bg);
  padding: 0.55rem 0.75rem;
  color: var(--kt-input-color);
  font-size: 0.925rem;
  border: 1px solid var(--kt-input-border-color);
  border-radius: 0.425rem;
  outline: 0 !important;
}

.select2-container--open .select2-dropdown {
  left: 0;
  top: -3px;
}

.select2-container .select2-dropdown .select2-search {
  padding: 0.5rem 1.25rem;
  margin: 0 0 0.5rem 0;
}

.select2-search--dropdown .select2-search__field {
  padding: 4px;
  width: 100%;
  box-sizing: border-box;
}

.select2-container .select2-search--dropdown {
  display: block;
  padding: 4px;
}

.select2-container--open .select2-dropdown--below {
  border-bottom: none;
  border-bottom-left-radius: 0;
  border-bottom-right-radius: 0;
}

.select2-container .select2-dropdown .select2-results__options {
  max-height: 250px;
  overflow-y: auto;
}

.select2-container .select2-results__options {
  list-style: none;
  margin: 0;
  padding: 0;
}

.select2-container .select2-dropdown .select2-results__option.select2-results__option--highlighted {
  background-color: var(--kt-component-hover-bg);
  color: var(--kt-component-hover-color);
  transition: color 0.2s ease;
}

.select2-container .select2-dropdown .select2-results__option {
  color: var(--kt-gray-700);
  transition: color 0.2s ease;
  padding: 0.75rem 1.25rem;
  margin: 0 0;
}

.select2-container .select2-results__option--selectable {
  cursor: pointer;
}

.select2-container .select2-results__option {
  padding: 6px;
  user-select: none;
  -webkit-user-select: none;
}
</style>