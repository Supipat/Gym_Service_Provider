import React from 'react';
import { Container } from 'react-bootstrap';

const Footer = () => {
  return (
    <footer className="py-4 bg-dark " style={{marginTop : "90px"}}>
      <Container className="text-center" style={{color:"#737373"}}>
        <h3 className="text-primary">GymService</h3>
        <p className="mb-0 text-secondary">© 2025 Team Notepads | Web Service.</p>
      </Container>
    </footer>
  );
};

export default Footer;