<script>
	import CommentBox from "../components/CommentBox.vue";

    export default {
		components: { CommentBox },
        data: function() {
            return {
                errormsg: null,
                loading: false,

				token: localStorage.getItem("token"),
				uname: localStorage.getItem("uname"),

				show_results: false,
				search_query: null,
				search_results: null,
				empty_results: true,

				empty_stream: true,
				stream: null,

				comments: {},

				show_likes: false,
				likes: null,

				modal: null,
            }
        },
        methods: {
            async search() {
				if (this.search_query !== "") {
					try {
						let response = await this.$axios.get("/user/" + this.uname + "/users?query_name=" + this.search_query, {
							headers: {
								Authorization: "Bearer " + this.token,
							}
						});

						this.search_results = response.data;

						this.show_results = true;

						if (this.search_results.users.length > 0) {
							this.empty_results = false;
						}
					} catch (e) {
                        if (e.response && e.response.status === 500) {
                            this.errormsg = "Something went wrong while trying to find results.";
                        } else {
                            this.errormsg = e.toString();
                        }
					}
				}
			},
			async getStream() {
				try {
					let response = await this.$axios.get("/user/" + this.uname + "/stream", {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

					this.stream = response.data;
					
					if (this.stream.photos.length > 0) {
						this.empty_stream = false;
					}
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to fetch the user's stream.";
					} else if (e.response && e.response.status == 401) {
						this.errormgs = "Forbidden access"
					} else {
						this.errormsg = e.toString();
					}
				}
			},
			async getPhotoComments(photo) {
				try {
					let response = await this.$axios.get("/user/" + photo.user.username + "/photos/" + photo.id + "/comments", {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

					this.comments = response.data;

					this.modal = new bootstrap.Modal(document.getElementById('logviewer'));
					this.modal.show();
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to fetch the user's stream.";
					} else if (e.response && e.response.status == 401) {
						this.errormgs = "Forbidden access"
					} else {
						this.errormsg = e.toString();
					}
				}
			},
			async updateLike(photo) {
				if (!photo.like_status) {
					try {
						let _ = await this.$axios.put("/user/" + photo.user.username + "/photos/" + photo.id + "/likes/" + this.uname, {}, {
							headers: {
								Authorization: "Bearer " + this.token,
							}
						});

						photo.like_count += 1;
					} catch (e) {
						if (e.response && e.response.status === 500) {
							this.errormsg = "Something went wrong while trying to register the like.";
						} else if (e.response && e.response.status == 401) {
							this.errormgs = "Forbidden access"
						} else {
							this.errormsg = e.toString();
						}
					}
				} else {
					try {
						let _ = await this.$axios.delete("/user/" + photo.user.username + "/photos/" + photo.id + "/likes/" + this.uname, {
							headers: {
								Authorization: "Bearer " + this.token,
							}
						});

						photo.like_count -= 1;
					} catch (e) {
						if (e.response && e.response.status === 500) {
							this.errormsg = "Something went wrong while trying to remove the like.";
						} else if (e.response && e.response.status == 401) {
							this.errormgs = "Forbidden access"
						} else {
							this.errormsg = e.toString();
						}
					}
				}

				photo.like_status = !photo.like_status;
			},
			async getPhotoLikes(photo) {
				try {
					let response = await this.$axios.get("/user/" + photo.user.username + "/photos/" + photo.id + "/likes", {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

					this.likes = response.data;

					this.show_likes = true;
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to retrieve likes.";
					} else if (e.response && e.response.status == 401) {
						this.errormgs = "Forbidden access"
					} else {
						this.errormsg = e.toString();
					}
				}
			},
			async logout() {
				localStorage.removeItem("token");
				localStorage.removeItem("uname");

                this.$router.push({path: "/"});
			},
			async profile() {
                this.$router.push({path: "/user/" + this.uname});
			}
        },
        mounted() {
			this.getStream();
		}
	}
</script>

<template>
    <div class="everything">
		<div class="header-div">
			<p class="header">Your stream</p>

			<div class="search-div">
				<input class="bar" placeholder="Search a profile" style="" v-model="search_query">

				<button class="button" @click="search" style="width: 9%; height: 70%;">
					<img class="button-image" src="/assets/search.svg"/>
				</button>
			</div>

			<div class="left-right-corner">
				<button @click="profile" class="button">
					<img class="user-icon" src="/assets/user-small.svg">
				</button>

				<button @click="logout" class="button">
					<img class="logout-icon" src="/assets/logout.svg">
				</button>
			</div>
        </div>

		<div v-if="!this.empty_stream" class="horizontal-scroll-panel">
			<div class="post-card" v-for="photo in this.stream.photos" :key="photo.id">
				<div class="post-card-header" style="margin-top: 0px">
					<RouterLink :to="'/user/' + photo.user.username" class="nav-link" style="margin-left: 20px; margin-top: 6px; height: 80px;">
						<p class="post-card-username">{{photo.user.username}}</p>
					</RouterLink>
				</div>

				<div class="post-photo-div">
					<div class="post-photo-bg"></div>
					<img class="post-photo-img" src="/assets/cupolone.jpg"> <!-- TODO: DA FARE -->
				</div>

				<div class="post-card-footer" style="margin-top: 7px">
					<button @click="updateLike(photo);" class="button" style="margin-bottom: 60px; margin-left: 20px">
						<div class="post-photo-utils" style="margin-right: 10px;">
							<img v-if="!photo.like_status" src="/assets/like-not-liked.svg"/>
							<img v-if="photo.like_status" src="/assets/like-liked.svg"/>
						</div>
					</button>

					<button @click="getPhotoLikes(photo)" class="button" style="margin: 5px 20px 0px 10px;">
						<p>{{photo.like_count}}</p>
					</button>

					<button @click="getPhotoComments(photo)" class="button" style="margin-bottom: 45px; margin-right: 8px">
						<div class="post-photo-utils">
							<img src="/assets/comment.svg"/>
						</div>
					</button>

					<p>{{photo.comment_count}}</p>

					<CommentBox id="logviewer" :comments="this.comments" :photo="photo" :modal="this.modal"></CommentBox>
				</div>
			</div>
		</div>

		<div v-if="this.empty_stream" class="horizontal-scroll-panel">
			<div style="display: flex; justify-content: center; margin-top: 13%">
				<p style="color: #485696; font-size: 300%">Find new users and follow your friends!</p>
			</div>
		</div>
    </div>

	<div v-if="this.show_likes" class="overlay">
		<div class="comment-box">
			<button class="button" @click="this.show_likes = false;" style="display:flex">
				<img class="cross" src="/assets/cross.svg"/>
			</button>
						
			<div class="search-scroll-panel">
				<div v-for="like in this.likes.users" :key="like.id">
					<div class="comment">
						<div class="comment-header">
							<div class="comment-op">
								<RouterLink @click="this.show_likes = false;" :to="'/user/' + like.username" class="nav-link">
									<p>{{like.username}}</p>
								</RouterLink>
							</div>
						</div>

						<div class="heightless-line"></div>
					</div>
				</div>
			</div>
		</div>
	</div>

	<div v-if="this.show_results" class="overlay">
		<div class="comment-box">
			<button class="button" @click="this.show_results = false;" style="display:flex">
				<img class="cross" src="/assets/cross.svg"/>
			</button>
						
			<div v-if="!this.empty_results" class="search-scroll-panel">
				<div v-for="result in this.search_results.users" :key="result.id">
					<div class="comment">
						<div class="comment-header">
							<div class="comment-op">
								<RouterLink @click="this.show_results = false;" :to="'/user/' + result.username" class="nav-link">
									<p>{{result.username}}</p>
								</RouterLink>
							</div>
						</div>

						<div class="heightless-line"></div>
					</div>
				</div>
			</div>

			<div v-if="this.empty_results" class="nothing-div">
				Nothing here!
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

	.nothing-div {
		display: flex;
		justify-content: center;

		margin-top: 28%;
		
		font-size: 200%;
		color: #485696;
	}

	.search-div {
		display: flex;
		justify-content: space-between;

		width: 37%;
		height: 20%;

		margin-top: 2%;
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
		width: 100%;
		/* border: solid 2px red; */

		margin-left: 3%;
		margin-top: -4%
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

        width: 83%;
		/* width: 100px; */
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