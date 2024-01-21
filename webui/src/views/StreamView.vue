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

			<div class="left-right-corner">
				<img class="user-icon" src="/assets/user-small.svg">
				<img class="logout-icon" src="/assets/logout.svg">
			</div>
        </div>

		<div class="horizontal-scroll-panel">
			<div class="post-card">
				<div class="post-card-header">
					<p class="post-card-username">walter4c</p>
				</div>

				<div class="post-photo-div">
					<div class="post-photo-bg"></div>
					<img class="post-photo-img" src="/assets/cupolone.jpg">
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

			<div class="post-card">
				<div class="post-card-header">
					<p class="post-card-username">enginemode1</p>
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

			<div class="post-card">
				<div class="post-card-header">
					<p class="post-card-username">enginemode1</p>
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

			<div class="post-card">
				<div class="post-card-header">
					<p class="post-card-username">enginemode1</p>
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

	<!-- <div class="overlay">
		<div class="comment-box">
			<img class="cross" src="/assets/cross.svg"/>

			<div class="comment-scroll-panel">
				<div class="comment">
					<div class="comment-header">
						<div class="comment-op">
							<p>enginemode1</p>
						</div>

						<img class="delete-comment" src="/assets/delete.svg"/>
					</div>

					<div class="comment-text">
						<p>lesgoooooooooooooooooooooooooooooooooooooooo ooooooooooooooooooooooooooooooooooooooooooooooooooooooo ooooooooooooooooooooooooooooooooooo</p>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-header">
						<div class="comment-op">
							<p>enginemode1</p>
						</div>
					</div>

					<div class="comment-text">
						<p>Lucian ammazzati</p>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-header">
						<div class="comment-op">
							<p>enginemode1</p>
						</div>
					</div>

					<div class="comment-text">
						<p>Lucian ammazzati</p>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-header">
						<div class="comment-op">
							<p>enginemode1</p>
						</div>
					</div>

					<div class="comment-text">
						<p>Lucian ammazzati</p>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-header">
						<div class="comment-op">
							<p>enginemode1</p>
						</div>
					</div>

					<div class="comment-text">
						<p>Lucian ammazzati</p>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-header">
						<div class="comment-op">
							<p>enginemode1</p>
						</div>
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
	</div> -->

	<div class="overlay">
		<div class="comment-box">
			<img class="cross" src="/assets/cross.svg"/>

			<div class="search-scroll-panel">
				<div class="comment">
					<div class="comment-header">
						<div class="comment-op">
							<p>enginemode1</p>
						</div>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-header">
						<div class="comment-op">
							<p>enginemode1</p>
						</div>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-header">
						<div class="comment-op">
							<p>enginemode1</p>
						</div>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-header">
						<div class="comment-op">
							<p>enginemode1</p>
						</div>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-header">
						<div class="comment-op">
							<p>enginemode1</p>
						</div>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-header">
						<div class="comment-op">
							<p>enginemode1</p>
						</div>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-header">
						<div class="comment-op">
							<p>enginemode1</p>
						</div>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-header">
						<div class="comment-op">
							<p>enginemode1</p>
						</div>
					</div>

					<div class="heightless-line"></div>
				</div>

				<div class="comment">
					<div class="comment-header">
						<div class="comment-op">
							<p>enginemode1</p>
						</div>
					</div>

					<div class="heightless-line"></div>
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

	.left-right-corner {
		display: flex;
		justify-content: space-between;

		width: 17%;
	}

	.comment-header {
		display: flex;
	}

	.delete-comment {
		width: 5%;
		height: 5%;

		margin-left: 3%;
		margin-top: -0.5%
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

	.search-scroll-panel {
		height: 100%;
		width: 97%;

		overflow: auto;
	}

	.comment-scroll-panel {
		height: 87%;
		width: 97%;

		overflow: auto;
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
</style>