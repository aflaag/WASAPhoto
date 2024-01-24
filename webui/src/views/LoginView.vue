<script setup>
	import ErrorMsg from "../components/ErrorMsg.vue";
</script>

<script>
    export default {
        components: { ErrorMsg },
        data: function() {
            return {
                errormsg: null,
                loading: false,

                username: "",

                user: null,
            }
        },
        methods: {
            async login() {
                if (this.username === "") {
                    this.errormsg = "The username can't be empty";
                } else {
                    try {
                        let response = await this.$axios.post("/session", {
                            username: this.username,
                        }, {});

                        this.user = response.data;

                        localStorage.setItem("token", this.user.id);
                        localStorage.setItem("uname", this.user.username);

                        this.errormsg = "";

                        this.$router.push({path: "/user/" + this.user.username + "/stream"});
                    } catch (e) {
                        if (e.response && e.response.status === 500) {
                            this.errormsg = "Something went wrong while trying to login.";
                        } else {
                            this.errormsg = e.toString();
                        }
                    }
                }
            },
        },
        mounted() {}
}
</script>

<template>
    <div class="everything">
        <div class="header-div">
            <p class="header">WASAPhoto</p>
        </div>

        <div class="center-div">
            <div class="user-icon-div">
                <img class="user-icon" src="/assets/user.svg">
            </div>

            <div class="bar-section-div">
                <input v-model="this.username" class="bar" placeholder="Enter your username!" style="text-align: center;"/>

                <button class="button" @click="login">
                    <img class="button-image" src="/assets/arrow.svg"/>
                </button>
            </div>
        </div>
    </div>

	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style>
    .everything {
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .header-div {
        margin-top: 1%;
        margin-left: 2%;

        width: 100%;
    }

    .header {
        font-weight: 800;
        font-size: 500%;
        color: #485696;
    }

    .center-div {
        display: flex;
        align-items: center;
        justify-content: center;

        position: absolute;

        top: 22%;
    }

    .user-icon-div {
        margin-bottom: 20%;
    }

    .bar-section-div {
        display: flex;
        align-items: center;
        justify-content: center;

        position: absolute;

        top: 100%;
        left: -80%;
    }

    .bar {
        background-color: #c6ddff;

        border-radius: 50px;
        border: 6px solid #485696;

        width: 600px;
        height: 60px;

        margin-right: 2%;

        font-size: 170%;
        color: #485696;

        box-sizing: border-box;
        padding: 0 30px 0 30px;
    }

    ::placeholder {
        color: #8a8a8a;

        opacity: 1;
    }

    .button {
        background: none;

        color: inherit;

        border: none;

        padding: 0;

        font: inherit;
        cursor: pointer;
        outline: inherit;
    }

    .button-image {
        width: 96%;
    }
</style>