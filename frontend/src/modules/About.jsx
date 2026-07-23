export default function About() {
	return (
		<section className="container my-2">
			<div className="box">
				<header className="has-b-border">
					<h2 className="is-size-4">About this website</h2>
				</header>
				<div className="mt-2">
					<p>It's a handmade site, the backend is made in <a href="https://go.dev" target="_blank">Go</a> and the frontend in <a href="https://react.dev" target="_blank">React</a> and <a href="https://bulma.io" target="_blank">Bulma</a> (who has the most ugly site in history...). It's hosted on oracle, with docker and stuff whatever nobody cares.</p>
					<span><small>All of the code for this site can be found <a href="https://github.com/Alvesafk/blog-2" target="_blank">here!</a></small></span>
				</div>
			</div>
		</section>
	);
}
