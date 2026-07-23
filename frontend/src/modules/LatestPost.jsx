import { useEffect, useState } from "react";
import { Link } from "react-router";
import Tags from "./Tags.jsx";

export default function LatestPost() {
	const [latestPost, setLatestPost] = useState(null);
	const [loading, setLoading] = useState(true);

	useEffect(() => {
		fetch(`/api/post/latest`)
			.then(response => {
				if (response.ok) {
					return response.json();
				}
				throw response;
			})
			.then(post => {
				setLatestPost(post);
				setLoading(false);
			})
			.catch(err => {
				console.error(`Error on fetching latest post: `, err);
				setLoading(false)
			});
	}, []);

	return (
		<section className="container column is-three-fifths">
			<div className="box box-fill-height">
				<header className="has-b-border">
					<h2 className="is-size-4">Updates</h2>
				</header>
				<div className="mt-2">
					{loading ? (
						<div class="skeleton-lines">
							<div></div>
							<div></div>
							<div></div>
							<div></div>
						</div>
					) : latestPost ? (
						<div className="latestPost">
							<h3>{latestPost.content.title}</h3>
							<p>{latestPost.content.preview}</p>
							<Link to={"/posts/" + latestPost.content.slugTitle} state={{ postId: latestPost.content.id }}>Read it!</Link> <br />
							<Tags tags={latestPost.content.tags} />
						</div>
					) : (
						<p>No post found!</p>
					)}
				</div>
			</div>
		</section>
	)
}
