import Header from '../modules/Header.jsx';
import Footer from '../modules/Footer.jsx';
import { Outlet } from 'react-router';

export default function MainLayout() {
	return (
		<>
			<Header />
			<main>
				<Outlet />
			</main>
			<Footer />
		</>
	);
}
