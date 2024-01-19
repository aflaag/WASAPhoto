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
            <p class="header">Your stream</p>

            <input class="bar" placeholder="Search a profile" style="margin-top: 2%;">

            <img class="user-icon" src="assets/user-small.svg">
        </div>

		<div class="horizontal-scroll-panel">
			<div class="post-card">
				<div class="post-card-header">
					<p class="post-card-username">walter4c</p>
				</div>

				<div class="post-photo-div">
					<div class="post-photo-bg"></div>
					<img class="post-photo-img" src="assets/cupolone.jpg">
				</div>

				<div class="post-card-footer">
					<div class="post-photo-utils">
						<img  src="assets/like-liked.svg">
					</div>

					<p>2.6k</p>

					<div class="post-photo-utils">
						<img src="assets/comment.svg">
					</div>

					<p>1.5k</p>
				</div>
			</div>

			<div class="post-card">
				<div class="post-card-header">
					<p class="post-card-username">enginemode1</p>
				</div>

				<div class="post-photo-div">
					<div class="post-photo-bg"></div>
					<img class="post-photo-img" src="assets/salisburgo.jpg">
				</div>

				<div class="post-card-footer">
					<div class="post-photo-utils">
						<img  src="assets/like-liked.svg">
					</div>

					<p>2.6k</p>

					<div class="post-photo-utils">
						<img src="assets/comment.svg">
					</div>

					<p>1.5k</p>
				</div>
			</div>

			<div class="post-card">
				<div class="post-card-header">
					<p class="post-card-username">enginemode1</p>
				</div>

				<div class="post-photo-div">
					<div class="post-photo-bg"></div>
					<img class="post-photo-img" src="assets/salisburgo.jpg">
				</div>

				<div class="post-card-footer">
					<div class="post-photo-utils">
						<img  src="assets/like-liked.svg">
					</div>

					<p>2.6k</p>

					<div class="post-photo-utils">
						<img src="assets/comment.svg">
					</div>

					<p>1.5k</p>
				</div>
			</div>

			<div class="post-card">
				<div class="post-card-header">
					<p class="post-card-username">enginemode1</p>
				</div>

				<div class="post-photo-div">
					<div class="post-photo-bg"></div>
					<img class="post-photo-img" src="assets/salisburgo.jpg">
				</div>

				<div class="post-card-footer">
					<div class="post-photo-utils">
						<img  src="assets/like-liked.svg">
					</div>

					<p>2.6k</p>

					<div class="post-photo-utils">
						<img src="assets/comment.svg">
					</div>

					<p>1.5k</p>
				</div>
			</div>
		</div>
    </div>

	<div class="overlay">
		<div class="comment-box">
			<img class="cross" src="assets/cross.svg"/>

			<div class="comment-scroll-panel">
				<div class="comment">
					<div class="comment-op">
						<p>enginemode1</p>
					</div>

					<div class="comment-text">
						<p>lesgoooooooooooooooooooooooooooooooooooooooo ooooooooooooooooooooooooooooooooooooooooooooooooooooooo ooooooooooooooooooooooooooooooooooo</p>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-op">
						<p>exyss_</p>
					</div>

					<div class="comment-text">
						<p>Lucian ammazzati</p>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-op">
						<p>exyss_</p>
					</div>

					<div class="comment-text">
						<p>Lucian ammazzati</p>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-op">
						<p>exyss_</p>
					</div>

					<div class="comment-text">
						<p>Lucian ammazzati</p>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-op">
						<p>exyss_</p>
					</div>

					<div class="comment-text">
						<p>Lucian ammazzati</p>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-op">
						<p>exyss_</p>
					</div>

					<div class="comment-text">
						<p>Lucian ammazzati</p>
					</div>

					<div class="heightless-line"></div>
				</div>
			</div>

			<div class="comment-input-box">
            	<input class="comment-bar" placeholder="Leave a comment!">
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

	.overlay {
		position: absolute;

		display: flex;
		justify-content: center;
		align-items: center;

		width: 100%;
		height: 100%;

		margin-top: -8.2%;
		
		background-color: rgba(0, 0, 0, 0.3);
	}

	.cross {
		position: absolute;

		width: 10%;
		height: 10%;

		z-index: 2;

		margin-top: 2%;
		margin-left: 85%;
	}

	.comment-bar {
        background-color: rgba(0, 0, 0, 0);

        border-radius: 0px;
        border: 0px solid rgba(0, 0, 0, 0);

        width: 93%;
		height: 50%;

        font-size: 110%;

		margin-left: 3.5%;
		/* margin-top: -40%; */

        box-sizing: boder-box;
	}

	.heightless-line {
		height: 2px;

		width: 99%;

		background-color: #757575;
	}

	.comment-scroll-panel {
		height: 87%;
		width: 97%;

		overflow: auto;
		/* white-space: nowrap; */
	}

	.comment-box {
		position: relative;

		width: 44%;
		height: 60%;

		border: solid 4px #485696;
		border-radius: 30px;

		background-color: #e7e7e7;

		font-size: 150%;

		z-index: 1;
	}

	.comment {
		margin-left: 4%;
		margin-top: 3%;

		/* border: solid 2px red; */
	}

	.comment-text {
		margin-top: -1%;

		font-weight: 300;
	}

    .header-div {
		display: flex;
        justify-content: space-between;

        margin-top: 1%;
        margin-left: 2%;
		margin-right: 2%;

        width: 100%;
    }

    .header {
        font-weight: 800;
        font-size: 500%;
        color: #485696;
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
		margin-top: 2%;
	}

	.post-card-header p {
		margin-left: 6%;
	}
</style>