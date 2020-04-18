import React from 'react';
import './App.css';
import Axios from 'axios';
import Loading from './loading/Loading';
import Empty from './empty/Empty';
import { Accordion, Card, Button } from 'react-bootstrap';

class App extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      transactions: [],
      showLoading: true
    }
  }

  componentDidMount = () => {
    Axios.get(`/api/v1/transactions`)
      .then(res => {
        this.setState({ transactions: res.data});
      })
      .catch(error => {
        console.error(error);
        this.setState({
          transactions: [],
        });
      })
      .finally(() => {
        this.setState({showLoading: false});
      });
  }

  renderTable = () => {
    return (
      <Accordion>
        {
          this.state.transactions.map(t => (
            <Card>
              <Card.Header>
                <Accordion.Toggle as={Button} variant="link" eventKey={t.id}>
                <div className={t.type === 'credit' ? "text-success" : "text-danger"}>
                  <p className={t.type === 'credit' ? '' : 'hidden'}> Credit transaction of +${t.ammount}</p>
                  <p className={t.type === 'debit' ? '' : 'hidden'}> Debit transaction of -${t.ammount}</p>
                </div>
                </Accordion.Toggle>
              </Card.Header>
              <Accordion.Collapse eventKey={t.id}>
                <Card.Body>
                  <div class="text-left text-monospace">
                    <p>ID: {t.id}</p>
                    <p>Type: {t.type}</p>
                    <p>Ammount: ${t.ammount}</p>
                    <p>Effective Date: {t.effective_date}</p>
                  </div>
                </Card.Body>
              </Accordion.Collapse>
            </Card>
          ))
        }
      </Accordion>
    )
  }

  render = () => {
    return (
      <div className="mt-3 ml-3 mr-3">
      <h1>Accounting Notebook</h1>
        <div className={this.state.showLoading ? '' : 'hidden'}>
          <Loading/>
        </div>
        <div className={(!this.state.showLoading && this.state.transactions.length === 0) ? '' : 'hidden'}>
          <Empty/>
        </div>
        <div className={(!this.state.showLoading && this.state.transactions.length > 0) ? '' : 'hidden'}>
          {this.renderTable()}
        </div>
      </div>
    )
  }
}

export default App;
