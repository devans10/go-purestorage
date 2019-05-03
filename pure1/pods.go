/* 
   Copyright 2018 David Evans

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package pure1

// PodService type creates a service to expose pods endpoints
type PodService struct {
	client *Client
}

// GetPods returns a list of pod objects
func (p *PodService) GetPods(params map[string]string) ([]Pod, error) {
	req, err := p.client.NewRequest("GET", "pods", params, nil)
	if err != nil {
		return nil, err
	}

	m := []Pod{}
	_, err = p.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
