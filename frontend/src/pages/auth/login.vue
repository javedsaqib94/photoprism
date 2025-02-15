<template>
  <v-container fluid fill-height class="auth-login wallpaper pa-3">
    <v-layout align-center justify-center>
      <v-flex xs12 sm8 md4 xl3 xxl2>
        <v-form ref="form" dense class="auth-login-form" accept-charset="UTF-8" @submit.prevent="login">
          <v-card class="elevation-12 auth-login-box blur-7">
            <v-card-text class="pa-3">
              <div class="logo text-xs-center">
                <img :src="$config.appIcon()" :alt="config.name">
              </div>
              <v-spacer></v-spacer>
              <v-text-field
                  v-model="username"
                  required hide-details solo flat light autofocus
                  type="text"
                  :disabled="loading"
                  name="username"
                  autocorrect="off"
                  autocapitalize="none"
                  :label="$gettext('Name')"
                  background-color="grey lighten-5"
                  class="input-name text-selectable"
                  :color="colors.accent"
                  :placeholder="$gettext('Name')"
                  prepend-icon="person"
              ></v-text-field>
              <v-text-field
                  v-model="password"
                  required hide-details solo flat light
                  :type="showPassword ? 'text' : 'password'"
                  :disabled="loading"
                  name="password"
                  autocorrect="off"
                  autocapitalize="none"
                  :label="$gettext('Password')"
                  background-color="grey lighten-5"
                  :placeholder="$gettext('Password')"
                  class="input-password mt-1 text-selectable"
                  :append-icon="showPassword ? 'visibility' : 'visibility_off'"
                  prepend-icon="lock"
                  :color="colors.accent"
                  @click:append="showPassword = !showPassword"
                  @keyup.enter.native="login"
              ></v-text-field>
              <v-spacer></v-spacer>
              <div class="pa-3 text-xs-center">
                <!-- a href="#" target="_blank" class="text-link px-2" :style="`color: ${colors.link}!important`"><translate>Forgot password?</translate></a -->
                <v-btn :color="colors.primary" round :disabled="loginDisabled"
                       class="white--text action-confirm px-3" @click.stop="login">
                  <translate>Sign in</translate>
                  <v-icon :right="!rtl" :left="rtl" dark>arrow_forward</v-icon>
                </v-btn>
              </div>
            </v-card-text>
          </v-card>
        </v-form>
      </v-flex>
    </v-layout>
    <footer>
      <v-layout wrap align-top pa-0 ma-0>
        <v-flex xs12 class="pa-0 body-2 text-selectable text-xs-center white--text" :class="[config.imprint ? 'text-sm-left sm6' : '']">
          <strong>{{ config.siteTitle }}</strong> – {{ config.siteCaption }}
        </v-flex>
        <v-flex v-if="config.imprint" xs12 sm6 class="pa-0 body-2 text-xs-center text-sm-right white--text">
          <a v-if="config.imprintUrl" :href="config.imprintUrl" target="_blank" class="text-link" :style="`color: ${colors.link}!important`">{{ config.imprint }}</a>
          <span v-else>{{ config.imprint }}</span>
        </v-flex>
      </v-layout>
    </footer>
  </v-container>
</template>

<script>

export default {
  name: "PPageAuthLogin",
  data() {
    const c = this.$config.values;

    return {
      colors: {
        accent: "#05dde1",
        primary: "#00adb0",
        link: "#c8e3e7",
      },
      loading: false,
      showPassword: false,
      username: "",
      password: "",
      sponsor: this.$config.isSponsor(),
      config: this.$config.values,
      siteDescription: c.siteDescription ? c.siteDescription : c.siteCaption,
      nextUrl: this.$route.params.nextUrl ? this.$route.params.nextUrl : "/",
      rtl: this.$rtl,
    };
  },
  computed: {
    loginDisabled() {
      return this.loading || !this.password || !this.username;
    }
  },
  created() {
    this.$scrollbar.hide(this.$isMobile);
  },
  destroyed() {
    this.$scrollbar.show();
  },
  methods: {
    login() {
      if (!this.username || !this.password) {
        return;
      }

      this.loading = true;
      this.$session.login(this.username, this.password).then(
        () => {
          this.loading = false;
          this.$router.push(this.nextUrl);
        }
      ).catch(() => this.loading = false);
    },
  }
};
</script>
