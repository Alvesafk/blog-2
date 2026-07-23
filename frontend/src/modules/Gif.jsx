export default function Gif({ url }) {
	return (
		<section className="container column">
			<div className="box">
				<header className="has-b-border">
					<h2 className="is-size-4">look at this guy...</h2>
				</header>
				<div className="mt-2">
					<figure className="image is-128x128 mx-auto">
						<img src={url} alt="SUPER GIF" />
					</figure>
				</div>
			</div>
		</section>
	)
}
