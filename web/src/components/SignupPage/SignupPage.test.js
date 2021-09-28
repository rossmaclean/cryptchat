import { render, screen } from '@testing-library/react';
import SignupPage from './SignupPage';

xtest('renders learn react link', () => {
  render(<SignupPage />);
  const linkElement = screen.getByText(/learn react/i);
  expect(linkElement).toBeInTheDocument();
});
