import React from 'react';
import './Loading.css'; 

class Loading extends React.Component {
    render = () => {
        return(
            <div className="Loading">
                <div className="d-flex justify-content-center">
                    <div className="spinner-border mt-5" role="status">
                        <span className="sr-only">Loading...</span>
                    </div>
                </div>
            </div>
        )
    }
}

export default Loading;