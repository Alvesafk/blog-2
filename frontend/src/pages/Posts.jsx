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
		<div className="post-list container">
			{loading ? (
				<p>Loading...</p>
			) : posts ? (
				<>
					{posts.content.map(post =>
						<PostCard key={post.id} post={post} />
					)}
				</>
			) : (
				<p>No post found!</p>
			)}
		</div>
	);
}

