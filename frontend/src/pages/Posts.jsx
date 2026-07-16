import { useState, useEffect } from 'react';
import PostCard from '../modules/PostCard';

export default function Posts() {
	const [posts, setPosts] = useState(null);
	const [loading, setLoading] = useState(true);

	useEffect(() => {
		fetch(`/api/post`)
			.then(response => {
				if (response.ok) {
					return response.json();
				}
				throw response;
			})
			.then(allPosts => {
				console.log(allPosts)
				setPosts(allPosts);
				setLoading(false);
			})
			.catch(err => {
				console.error(`Error on fetching latest post: `, err);
				setLoading(false)
			});
	}, []);

	return (
		<>
			<h2>This is the Posts</h2>
			{loading ? (
				<p>Loading...</p>
			) : posts ? (
				<div className="post-list">
					{posts.content.map(post =>
						<PostCard post={post} />
					)}
				</div>
			) : (
				<p>No post found!</p>
			)}
		</>
	);
}

