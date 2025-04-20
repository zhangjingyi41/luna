<template>
    <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" :ellipsis="false"
        @select="handleSelect">
        <el-menu-item index="0">
            <el-icon><HomeFilled /></el-icon>
        </el-menu-item>
        <el-menu-item index="1" @click="handleLogin">登录</el-menu-item>
    </el-menu>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import http from './utils/axios'
import {generateRandomCode} from './utils/tools'

const activeIndex = ref('0')

// 处理登录按钮点击事件
const handleLogin = async () => {
    try {
        // const response = await http.get('oauth2/oauthorization', {
        //     params: {
        //         client_id: '6FF9C917-3B38-7311-F146-4B617EF899BD',
        //         redirect_uri: 'http://localhost:5174',
        //         response_type: 'code',
        //         scope: 'default',
        //         state: generateRandomCode(16),
        //     }
        // })
        // console.log('登录成功:', response.data)
        // // 使用 response.data.data.target_url的地址，打开新的标签页
        // if (response.data && response.data.data && response.data.data.target_url) {
        //     const data = response.data.data
        //     window.open(`${data.target_url}?client_id=${data.client_id}&redirect_uri=${data.redirect_uri}&response_type=${data.response_type}&scope=${data.scope}&state=${data.state}`, '_blank')
        // } else {
        //     console.error('未找到 target_url')
        // }
        const client_config = {
            client_id: '6FF9C917-3B38-7311-F146-4B617EF899BD',
            redirect_uri: 'http://localhost:5174',
            response_type: 'code',
            scope: 'default',
            state: generateRandomCode(16),
        }
        window.open(`http://localhost:5173?client_id=${client_config.client_id}&redirect_uri=${client_config.redirect_uri}&response_type=${client_config.response_type}&scope=${client_config.scope}&state=${client_config.state}`)
    } catch (error) {
        console.error('登录失败:', error)
    }
}

const handleSelect = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
</script>

<style scoped>
.el-menu--horizontal>.el-menu-item:nth-child(1) {
    margin-right: auto;
}
/* 修改子菜单字体大小 */
::v-deep(.el-sub-menu__title) {
    font-size: 18px;
}

::v-deep(.el-menu-item) {
    font-size: 18px;
}
</style>
