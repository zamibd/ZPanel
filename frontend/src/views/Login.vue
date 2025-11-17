<template>
  <v-container class="fill-height" style="margin-top: 100px;">
    <v-row justify="center" align="center">
      <v-col cols="12" sm="8" md="4">
        <v-card>
          <v-card-title class="headline" v-text="$t('login.title')"></v-card-title>
          <v-card-text>
            <v-form @submit.prevent="login" ref="form">
              <v-text-field v-model="username" prepend-icon="mdi-account" :label="$t('login.username')" :rules="usernameRules" required></v-text-field>
              <v-text-field v-model="password" prepend-icon="mdi-form-textbox-password" :label="$t('login.password')" :rules="passwordRules" type="password" required></v-text-field>
      <v-text-field v-model="passcode" prepend-icon="mdi-shield-lock-outline" :label="$t('login.passcode')" :rules="passcodeRules" required></v-text-field>
              <v-btn :loading="loading" type="submit" color="primary" block class="mt-2" v-text="$t('actions.submit')"></v-btn>
            </v-form>
            <v-select
              density="compact"
              class="mt-2"
              hide-details
              variant="solo"
              :items="languages"
              v-model="$i18n.locale"
              @update:modelValue="changeLocale">
              <template v-slot:append>
                <v-menu>
                  <template v-slot:activator="{ props }">
                    <v-btn icon v-bind="props">
                      <v-icon>mdi-theme-light-dark</v-icon>
                    </v-btn>
                  </template>
                  <v-list>
                    <v-list-item
                      v-for="th in themes"
                      :key="th.value"
                      @click="changeTheme(th.value)"
                      :prepend-icon="th.icon"
                      :active="isActiveTheme(th.value)"
                    >
                      <v-list-item-title>{{ $t(`theme.${th.value}`) }}</v-list-item-title>
                    </v-list-item>
                  </v-list>
                </v-menu>
              </template>
            </v-select>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts" setup>
import { ref, computed } from "vue"
import { useLocale,useTheme } from 'vuetify'
import { i18n, languages } from '@/locales'
import { useRouter } from 'vue-router'
import HttpUtil from '@/plugins/httputil'


const theme = useTheme()
const locale = useLocale()

const themes = [
{ value: 'light', icon: 'mdi-white-balance-sunny' },
{ value: 'dark', icon: 'mdi-moon-waning-crescent' },
{ value: 'system', icon: 'mdi-laptop' },
]

const username = ref('')
const usernameRules = [
(value: string) => {
  if (value?.length > 0) return true
  return i18n.global.t('login.unRules')
},
]

const password = ref('')
const passwordRules = [
(value: string) => {
  if (value?.length > 0) return true
  return i18n.global.t('login.pwRules')
},
]

const passcode = ref('')
const passcodeRules = [
(value: string) => {
  if (value?.length == 6) return true
  return i18n.global.t('login.passcodeRules')
},
]

const loading = ref(false)
const router = useRouter()

const login = async () => {
if (username.value == '' || password.value == '' || passcode.value == '') return
loading.value=true
const response = await HttpUtil.post('api/login',{user: username.value, pass: password.value, passcode: passcode.value})
if(response.success){
  setTimeout(() => {
    loading.value=false
    router.push('/')
  }, 500)
} else {
  loading.value=false
}
}
const changeLocale = (l: any) => {
locale.current.value = l ?? 'en'
localStorage.setItem('locale', locale.current.value)
}
const changeTheme = (th: string) => {
theme.change(th)
localStorage.setItem('theme', th)
}
const isActiveTheme = (th: string) => {
const current = localStorage.getItem('theme') ?? 'system'
return current == th
}

const selectedTheme = computed({
get: () => localStorage.getItem('theme') ?? 'system',
set: () => {}
})
</script>

<style scoped>
.minimal-login {
min-height: 100vh;
background: #0a0a0a;
position: relative;
overflow: hidden;
}

.particles-bg {
position: absolute;
top: 0;
left: 0;
width: 100%;
height: 100%;
background-image: 
  radial-gradient(2px 2px at 20px 30px, #ffffff10, transparent),
  radial-gradient(2px 2px at 40px 70px, #ffffff08, transparent),
  radial-gradient(1px 1px at 90px 40px, #ffffff15, transparent),
  radial-gradient(1px 1px at 130px 80px, #ffffff10, transparent);
background-repeat: repeat;
background-size: 200px 200px;
animation: particleMove 20s linear infinite;
}

@keyframes particleMove {
0% { transform: translate(0, 0); }
100% { transform: translate(-200px, -200px); }
}

.login-container {
position: relative;
z-index: 1;
}

.logo-wrapper {
animation: logoFloat 3s ease-in-out infinite;
}

@keyframes logoFloat {
0%, 100% { transform: translateY(0px); }
50% { transform: translateY(-10px); }
}

.login-header h1 {
color: white;
letter-spacing: 3px;
}

.input-group {
position: relative;
}

.input-label {
display: block;
color: #ffffff80;
font-size: 0.875rem;
font-weight: 500;
margin-bottom: 8px;
text-transform: uppercase;
letter-spacing: 1px;
}

.minimal-input :deep(.v-field) {
background: rgba(255, 255, 255, 0.05) !important;
border: 1px solid rgba(255, 255, 255, 0.1);
border-radius: 4px !important;
transition: all 0.3s ease;
}

.minimal-input :deep(.v-field):hover {
background: rgba(255, 255, 255, 0.08) !important;
border-color: rgba(255, 255, 255, 0.2);
}

.minimal-input :deep(.v-field--focused) {
background: rgba(255, 255, 255, 0.1) !important;
border-color: rgb(var(--v-theme-primary)) !important;
box-shadow: 0 0 0 1px rgb(var(--v-theme-primary));
}

.login-submit-btn {
background: linear-gradient(45deg, rgb(var(--v-theme-primary)), rgb(var(--v-theme-secondary))) !important;
text-transform: none;
font-weight: 600;
letter-spacing: 1px;
transition: all 0.3s ease;
}

.login-submit-btn:hover {
transform: translateY(-2px);
box-shadow: 0 8px 25px rgba(var(--v-theme-primary), 0.3);
}

.footer-controls {
gap: 16px;
}

/* Light theme adjustments */
.v-theme--light .minimal-login {
background: #f8f9fa;
}

.v-theme--light .particles-bg {
background-image: 
  radial-gradient(2px 2px at 20px 30px, #00000010, transparent),
  radial-gradient(2px 2px at 40px 70px, #00000008, transparent),
  radial-gradient(1px 1px at 90px 40px, #00000015, transparent),
  radial-gradient(1px 1px at 130px 80px, #00000010, transparent);
}

.v-theme--light .login-header h1 {
color: #1a1a1a;
}

.v-theme--light .input-label {
color: #666;
}

.v-theme--light .minimal-input :deep(.v-field) {
background: rgba(0, 0, 0, 0.02) !important;
border-color: rgba(0, 0, 0, 0.1);
}

.v-theme--light .minimal-input :deep(.v-field):hover {
background: rgba(0, 0, 0, 0.04) !important;
border-color: rgba(0, 0, 0, 0.2);
}

/* Mobile-specific styles */
@media (max-width: 600px) {
.login-container {
  padding: 0 16px;
}

.login-card {
  padding: 24px 20px !important;
}

.particles-bg {
  /* Reduce animation intensity on mobile for better performance */
  animation-duration: 30s;
  opacity: 0.7;
}

.mobile-friendly-input :deep(.v-field) {
  font-size: 16px; /* Prevents zoom on iOS */
  min-height: 48px; /* Touch-friendly height */
}

.mobile-friendly-input :deep(.v-field__input) {
  padding: 12px 16px;
}

.mobile-friendly-btn {
  font-size: 16px;
  min-height: 48px; /* Touch-friendly height */
  padding: 12px 16px;
}

.mobile-friendly-select :deep(.v-select__selections) {
  font-size: 14px;
}

.mobile-friendly-select :deep(.v-field) {
  min-height: 40px;
}

/* Responsive typography */
.responsive-title {
  font-size: 1.5rem !important;
  line-height: 1.3;
}

.responsive-subtitle {
  font-size: 0.875rem !important;
}

/* Better touch targets for theme toggle */
.footer-controls .v-btn-toggle .v-btn {
  min-width: 44px;
  min-height: 44px;
}

/* Reduce margins on mobile */
.input-group {
  margin-bottom: 20px !important;
}

.login-submit-btn {
  margin-bottom: 24px !important;
}
}

/* Tablet styles */
@media (min-width: 601px) and (max-width: 960px) {
.responsive-title {
  font-size: 1.75rem !important;
}

.responsive-subtitle {
  font-size: 1rem !important;
}
}

/* Desktop styles */
@media (min-width: 961px) {
.responsive-title {
  font-size: 2rem !important;
}

.responsive-subtitle {
  font-size: 1.125rem !important;
}
}

/* High DPI displays */
@media (-webkit-min-device-pixel-ratio: 2), (min-resolution: 192dpi) {
.responsive-logo {
  /* Ensure crisp icons on retina displays */
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
}
</style>
