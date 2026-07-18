import { useState, useEffect } from 'react'

export default function Currently() {
	const [currently, setCurrently] = useState(null);
	const [loading, setLoading] = useState(true);

	useEffect(() => {
		fetch(`/api/currently`)
			.then(response => {
				if (response.ok) {
					return response.json();
				}
				throw response;
			})
			.then(c => {
				setCurrently(c);
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
			) : currently ? (
				<div className="latestPost">
					<h2>Currently...</h2>
					<p>{currently.content.content}</p>
					<p><small>Last updated at {new Date(currently.content.lastUpdatedAt).toLocaleString()}!</small></p>
				</div>
			) : (
				<p>No post found!</p>
			)}
		</div>
	)
}
