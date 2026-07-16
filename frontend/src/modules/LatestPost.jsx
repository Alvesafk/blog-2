import { useEffect, useState } from "react";
import { Link } from "react-router";

export default function LatestPost() {
	function printTags(tags) {
		let result = "";
		tags.forEach(function(element, index, array) {
			if (index == array.length - 1) {
				result += element;
			} else {
				result += element + ", ";
			}
		});

		return (
			<span>
				<small>
					{result}
				</small>
			</span>
		);
	}

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
		<div>
			<h2>Updates</h2>
			{loading ? (
				<p>Loading...</p>
			) : latestPost ? (
				<div className="latestPost">
					<h3>{latestPost.content.title}</h3>
					<p>{latestPost.content.preview}</p>
					<Link to={"/posts/" + latestPost.content.slug_title}>Read it!</Link> <br/>
					{printTags(latestPost.content.tags)}
				</div>
			) : (
				<p>No post found!</p>
			)}
		</div>
	)
}
