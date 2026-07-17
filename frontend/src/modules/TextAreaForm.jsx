import { useState } from 'react'

export default function TextAreaForm({ url, onCommentPosted }) {
	const [formData, setFormData] = useState({
		author: "",
		content: "",
	})
	const [loading, setLoading] = useState(false)
	const [error, setError] = useState(null)

	function handleChange(e) {
		const { name, value } = e.target
		setFormData((prev) => ({ ...prev, [name]: value }))
	}

	async function handleSubmit(e) {
		e.preventDefault()
		setLoading(true)
		setError(null)

		try {
			const response = await fetch(url, {
				method: "POST",
				headers: {
					"Content-Type": "application/json"
				},
				body: JSON.stringify(formData)
			})

			if (!response.ok) {
				throw new Error(`Error on post: ${response.status}`)
			}

			onCommentPosted()
		} catch (err) {
			setError(err.message)
		} finally {
			setFormData({
				author: "",
				content: "",
			})
			setLoading(false)
		}
	}

	return (
		<form onSubmit={handleSubmit}>
			<input
				type="text"
				name="author"
				placeholder='Who are you?'
				value={formData.author}
				onChange={handleChange}
			/>
			<textarea
				rows="3"
				cols="100"
				placeholder='Write a comment!'
				name='content'
				value={formData.content}
				onChange={handleChange}
			></textarea>
			<button type="submit" disabled={loading}>
				{loading ? "Posting..." : "Post your comment!"}
			</button>
			{error && <p>{error}</p>}
		</form>
	)
}
