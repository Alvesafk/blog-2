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
		<section className="latestPost column is-two-fifths container ">
			<div className='box box-fill-height'>
				<header className='has-b-border'>
					<h2 className="is-size-4">Currently...</h2>
				</header>
				<div className='mt-2'>
					{loading ? (
						<p>Loading...</p>
					) : currently ? (
						<>
							<p>{currently.content.content}</p>
							<p><small className='is-size-7'>Last updated at {new Date(currently.content.lastUpdatedAt).toLocaleString()}!</small></p>
						</>
					) : (
						<p>Doing nothing?</p>
					)
					}
				</div>
			</div>
		</section >
	)
}
