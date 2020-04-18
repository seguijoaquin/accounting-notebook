import React from 'react';
import './Empty.css';

class Empty extends React.Component {
    render = () => {
        return (
            <div className="Empty">
                <p>No transactions.</p>
            </div>
        )
    }
}

export default Empty;