<script>
    export default {
        data: function() {
            return {
                errormsg: null,
                loading: false,
                some_data: null,
            }
        },
        methods: {
            async refresh() {
                this.loading = true;
                this.errormsg = null;
                try {
                    let response = await this.$axios.get("/stream"); // TODO: DA CAMBIARE
                    this.some_data = response.data;
                } catch (e) {
                    this.errormsg = e.toString();
                }
                this.loading = false;
            },
        },
        mounted() {
            this.refresh()
        }
}
</script>

<template>
    <div class="everything">
        <div class="header-div">
            <div class="username-div">
                <div style="display: flex; justify-content: center;">
                    <p class="header">a.flag_</p>
                </div>

                <div class="option-buttons-div">
                    <img class="follow-icon" src="/assets/post.svg">

                    <img class="ban-icon" src="/assets/username.svg">
                </div>
            </div>

            <div class="numbers-div">
                <div class="numbers-block-div">
                    <p class="numbers-block-numbers">52</p>
                    <p>Photos</p>
                </div>

                <div class="numbers-block-div">
                    <p class="numbers-block-numbers">0</p>
                    <p style="margin-top: 3%">Followers</p>
                </div>

                <div class="numbers-block-div">
                    <p class="numbers-block-numbers">1504</p>
                    <p style="margin-top: 3%">Following</p>
                </div>
            </div>

            <img class="user-icon" src="/assets/home.svg">
        </div>

		<div class="horizontal-scroll-panel">
			<div class="post-card">
				<div class="post-card-header">
					<p class="post-card-username">walter4c</p>

					<img class="delete-photo" src="/assets/delete.svg"/>
				</div>

				<div class="post-photo-div">
					<div class="post-photo-bg"></div>
					<img class="post-photo-img" src="/assets/cupolone.jpg">
				</div>

				<div class="post-card-footer">
					<div class="post-photo-utils">
						<img src="/assets/like-liked.svg">
					</div>

					<p>2.6k</p>

					<div class="post-photo-utils">
						<img src="/assets/comment.svg">
					</div>

					<p>1.5k</p>
				</div>
			</div>

			<div class="post-card">
				<div class="post-card-header">
					<p class="post-card-username">walter4c</p>
				</div>

				<div class="post-photo-div">
					<div class="post-photo-bg"></div>
					<img class="post-photo-img" src="/assets/salisburgo.jpg">
				</div>

				<div class="post-card-footer">
					<div class="post-photo-utils">
						<img  src="/assets/like-liked.svg">
					</div>

					<p>2.6k</p>

					<div class="post-photo-utils">
						<img src="/assets/comment.svg">
					</div>

					<p>1.5k</p>
				</div>
			</div>
		</div>
    </div>
</template>

<style>
    .everything {
        display: flex;
        justify-content: center;
        align-items: center;
    }

	.delete-photo {
		width: 10%;
		height: 10%;

		margin-top: 3%;
	}

    .numbers-div {
        display: flex;
        justify-content: space-between;

        margin-top: 2%;

        width: 33%;

        font-weight: 200;
        font-size: 50px;

        color: #485696;
    }

    .numbers-block-div {
        display: flex;
        /* justify-content: center; */
        text-align: center;
        flex-direction: column;
    }

    .numbers-block-numbers {
        font-weight: 600;
        margin-bottom: -10%
    }

    .username-div {
        width: 41%;
    }

    .header {
        font-weight: 800;
        font-size: 500%;
        color: #485696;
    }

    .option-buttons-div {
        /* width: 60%; */
        display: flex;
        justify-content: space-between;

		/* border: solid 3px green; */
    }

	.horizontal-scroll-panel {
		position: absolute;

		/* border: solid 3px red; */

		top: 30%;
		left: 2%;

		width: 96%;
		height: 60%;

		overflow: auto;
		white-space: nowrap;
	}

	.post-card {
		display: inline-block;
		text-align: center;

		width: 25%;
		height: 100%;

		margin-right: 3%;

		border: solid 4px #485696;
		border-radius: 30px;

        color: #485696;

		font-size: 300%;
	}

	.post-photo-div {
		width: 100%;
		height: 70%;

		margin-top: -2%;

		/* border: solid 3px blue; */

		z-index: 1;
	}

	.post-photo-bg {
		width: 100%;
		height: 100%;
		/* object-fit: fill;
		overflow: hidden; */
		background-color: #c8cfe8;

		z-index: 2;
	}

	.post-photo-img {
		width: 100%;
		height: 100%;

		object-fit: contain;
		overflow: hidden;

		z-index: 3;

		/* border: solid 3px blue; */

		/* margin-left: -100%; */
		margin-top: -100%;
	}

	.post-card-footer {
		display: flex;
		font-weight: 800;

		/* border: solid 1px green; */
		height: 14%;
	}

	.post-photo-utils {
		/* border: solid 1px green; */
		height: 90%;

		margin-left: 3%;
	}

	.post-card-footer p {
		margin-top: 1.5%;
		margin-right: 4%;
		margin-left: 2%;
	}

	.post-card-header {
		display: flex;
		justify-content: space-between;
		margin-top: 2%;

		width: 94%;
	}

	.post-card-header p {
		margin-left: 6%;
	}
</style>