import { Link } from "react-router";

export default function Header() {
	return (
		<header>
			<nav className="navbar">
				<div className="navbar-brand">
					<h1 className="navbar-item title">
						<Link className="has-text-primary-00" to="/" prefetch="intent">Alves's Blog</Link>
					</h1>
				</div>
				<div className="navbar-menu">
					<div className="navbar-end">
						<Link className="navbar-item" to="/" prefetch="intent">Home</Link>
						<Link className="navbar-item" to="/posts" prefetch="intent">Posts</Link>
						<Link className="navbar-item" to="/resources">Resources</Link>
						<Link className="navbar-item" to="/guestbook" prefetch="intent">Guestbook</Link>
					</div>
				</div>
			</nav>
		</header>
	);
}
